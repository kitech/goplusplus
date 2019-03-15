package gopp

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

func LoadTLSCertKeyFromOneFile(certkeyfile string) (tls.Certificate, error) {
	bcc, err := ioutil.ReadFile(certkeyfile)
	ErrPrint(err)
	if err != nil {
		return tls.Certificate{}, err
	}

	return LoadTLSCertKeyFromOneBuffer(bcc)
}

func LoadTLSCertKeyFromTwoFile(certFile, keyFile string) (tls.Certificate, error) {
	// certFile, keyFile := filepath.Join(dir, "cert.pem"), filepath.Join(dir, "key.pem")
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	ErrPrint(err)
	return cert, err
}

func LoadTLSCertKeyFromOneBuffer(bcc []byte) (tls.Certificate, error) {
	// top: private key part
	// bottom: public cert part
	certparts := strings.Split(string(bcc), "-----BEGIN CERTIFICATE-----")
	certPEMBlock := []byte("-----BEGIN CERTIFICATE-----" + certparts[1])
	keyPEMBlock := []byte(certparts[0])
	certo, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	ErrPrint(err)
	return certo, err
}

func LoadTLSCertKeyFromTwoBuffer(certbcc, keybcc []byte) (tls.Certificate, error) {
	// top: private key part
	// bottom: public cert part
	certPEMBlock := certbcc
	keyPEMBlock := keybcc
	certo, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	ErrPrint(err)
	return certo, err
}

/*
openssl x509 -noout -fingerprint -sha256 -inform pem -in [certificate-file.crt]
openssl x509 -noout -fingerprint -sha1 -inform pem -in [certificate-file.crt]
openssl x509 -noout -fingerprint -md5 -inform pem -in [certificate-file.crt]
*/
func TLSCertKeyFPSha1(cert tls.Certificate) string {
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	ErrPrint(err)
	return Sha1AsStr(x509Cert.Raw)
}
func TLSCertKeyFPSha256(cert tls.Certificate) string {
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	ErrPrint(err)
	return Sha256AsStr(x509Cert.Raw)
}
func TLSCertKeyFPMd5(cert tls.Certificate) string {
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	ErrPrint(err)
	return Md5AsStr(x509Cert.Raw)
}

func SaveTLSCertKeyOneFile(cert tls.Certificate, fname string) error {
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	ErrPrint(err)

	if false {
		log.Println("Certn", len(cert.Certificate))
		log.Println("Leaf", cert.Leaf != nil)
		log.Println("OCSPStaple", cert.OCSPStaple != nil)
		log.Println("PrivateKey", cert.PrivateKey != nil)
		log.Println("type", reflect.TypeOf(cert.PrivateKey))

		pko := cert.PrivateKey.(*rsa.PrivateKey)
		pkov := reflect.ValueOf(pko)
		szmthv := pkov.MethodByName("Size")
		if !szmthv.IsNil() { // fix 1.11+ method ras.PrivateKey.Size()
			log.Println("pksz:", szmthv.Call(nil), pko.Validate())
			// log.Println("pksz:", pko.Size(), pko.Validate())
		}
	}

	derBytes := x509Cert.Raw
	priv := cert.PrivateKey

	certbuf := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keybuf := pem.EncodeToMemory(pemBlockForKey(priv))

	if false {
		// keybuf := []byte{}
		log.Println("certlen:", len(certbuf))
		log.Println("privkeylen:", len(keybuf))

		privblk, _ := pem.Decode(keybuf)
		log.Println(privblk.Type, privblk.Headers, len(privblk.Bytes), Md5AsStr(privblk.Bytes))
	}

	mrgbuf := append(keybuf, certbuf...)
	return ioutil.WriteFile(fname, mrgbuf, 0644)
}
func SaveTLSCertKeyTwoFile(cert tls.Certificate, certFile, keyFile string) error {
	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	ErrPrint(err)

	derBytes := x509Cert.Raw
	priv := cert.PrivateKey

	certbuf := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keybuf := pem.EncodeToMemory(pemBlockForKey(priv))

	err = ioutil.WriteFile(certFile, certbuf, 0644)
	if err == nil {
		err = ioutil.WriteFile(keyFile, keybuf, 0644)
	}
	return err
}

//
var _rdr = rand.Reader
var _x509ct = x509.Certificate{}

func NewTLSCertificate(hosts []string, days int, isCA bool, rsaBits int) (tls.Certificate, error) {
	certbuf, keybuf, err := NewX509Certificate(hosts, days, isCA, rsaBits)
	ErrPrint(err)

	return tls.X509KeyPair(certbuf, keybuf)
}

func NewX509CertToOneFile(hosts []string, days int, isCA bool, rsaBits int, pemFile string) error {
	pembuf, err := NewX509CertOne(hosts, days, isCA, rsaBits)
	ErrPrint(err)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(pemFile, pembuf, 0644)
	return err
}
func NewX509CertToTwoFile(hosts []string, days int, isCA bool, rsaBits int, certFile, keyFile string) error {
	certbuf, keybuf, err := NewX509Certificate(hosts, days, isCA, rsaBits)
	ErrPrint(err)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(certFile, certbuf, 0644)
	if err == nil {
		err = ioutil.WriteFile(keyFile, keybuf, 0644)
	}
	return err
}
func NewX509CertOne(hosts []string, days int, isCA bool, rsaBits int) ([]byte, error) {
	certbuf, keybuf, err := NewX509Certificate(hosts, days, isCA, rsaBits)
	ErrPrint(err)
	mrgbuf := append(keybuf, certbuf...)
	return mrgbuf, err
}

/*
hosts: default nil
days: must > 0
rsaBits: default 0
*/
func NewX509Certificate(hosts []string, days int, isCA bool, rsaBits int) ([]byte, []byte, error) {
	host := strings.Join(hosts, ",")
	validFor := time.Duration(days) * 24 * time.Hour
	if rsaBits <= 0 {
		rsaBits = 2048
	}
	validFrom := ""
	ecdsaCurve := ""
	certbuf, keybuf, err := main_tls_generate_cert(&host, &validFrom, &validFor, &isCA, &rsaBits, &ecdsaCurve)
	ErrPrint(err)

	return certbuf, keybuf, err
}

/*
var (
	host = flag.String("host", "", "Comma-separated hostnames and IPs to generate a certificate for")
	validFrom = flag.String("start-date", "", "Creation date formatted as Jan 1 15:04:05 2011")
	validFor = flag.Duration("duration", 365*24*time.Hour, "Duration that certificate is valid for")
	isCA = flag.Bool("ca", false, "whether this cert should be its own Certificate Authority")
	rsaBits = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set")
	ecdsaCurve = flag.String("ecdsa-curve", "", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
)
*/

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

func main_tls_generate_cert(host *string, validFrom *string, validFor *time.Duration, isCA *bool, rsaBits *int, ecdsaCurve *string) ([]byte, []byte, error) {
	// flag.Parse()
	if len(*host) == 0 {
		// log.Fatalf("Missing required --host parameter")
	}

	var priv interface{}
	var err error

	switch *ecdsaCurve {
	case "":
		priv, err = rsa.GenerateKey(rand.Reader, *rsaBits)
	case "P224":
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized elliptic curve: %q", *ecdsaCurve)
		os.Exit(1)
	}

	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	var notBefore time.Time
	if len(*validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", *validFrom)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse creation date: %s\n", err)
			os.Exit(1)
		}
	}

	notAfter := notBefore.Add(*validFor)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},

		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := strings.Split(*host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}
	template.Issuer = pkix.Name{Organization: []string{*host}, CommonName: *host}
	template.Subject = pkix.Name{Organization: []string{*host}, CommonName: *host}

	if *isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	certbuf := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keybuf := pem.EncodeToMemory(pemBlockForKey(priv))
	return certbuf, keybuf, err

	/*
		certOut, err := os.Create("cert.pem")
		if err != nil {
			log.Fatalf("failed to open cert.pem for writing: %s", err)
		}

		if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
			log.Fatalf("failed to write data to cert.pem: %s", err)
		}

		if err := certOut.Close(); err != nil {
			log.Fatalf("error closing cert.pem: %s", err)
		}

		log.Print("wrote cert.pem\n")
		keyOut, err := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatalf("failed to open key.pem for writing: %s", err)
		}

		if err := pem.Encode(keyOut, pemBlockForKey(priv)); err != nil {
			log.Fatalf("failed to write data to key.pem: %s", err)
		}

		if err := keyOut.Close(); err != nil {
			log.Fatalf("error closing key.pem: %s", err)
		}

		log.Print("wrote key.pem\n")
	*/
}

// https://golang.org/src/crypto/tls/generate_cert.go
// x509.CreateCertificate(rand io.Reader, template *x509.Certificate, parent *x509.Certificate, pub interface{}, priv interface{})

// Setup a bare-bones TLS config for the server
func GenerateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
}
