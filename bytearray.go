package gopp

import "bytes"

// aims to clone qt's QByteArray methods
type _ByteArray struct{ dptr *[]byte }
type ByteArray *_ByteArray

func NewByteArray(buf ...[]byte) ByteArray {
	if len(buf) == 0 {
		return &_ByteArray{&[]byte{}}
	} else {
		tmp := BytesDup(buf[0])
		return &_ByteArray{&tmp}
	}
}
func NewByteArrayNChars(size int, ch byte) ByteArray {
	tmp := bytes.Repeat([]byte{ch}, size)
	return &_ByteArray{&tmp}
}

// if not n, then full
func NewByteArrayString(s string, n ...int) ByteArray {
	if len(n) > 0 {
		tmp := []byte(StrSuf(s, n[0]))
		return &_ByteArray{&tmp}
	} else {
		tmp := []byte(s)
		return &_ByteArray{&tmp}
	}
}

func NewByteArrayOther(other ByteArray) ByteArray { return NewByteArray(*other.dptr) }
func (this *_ByteArray) Clone() ByteArray         { return NewByteArray(*this.dptr) }
func (this *_ByteArray) AppendOther(other ByteArray) ByteArray {
	tmp := append(*this.dptr, *other.dptr...)
	this.dptr = &tmp
	return this
}

func (this *_ByteArray) AppendString(s string, n ...int) ByteArray {
	if len(n) > 0 {
		tmp := append(*this.dptr, []byte(StrSuf(s, n[0]))...)
		this.dptr = &tmp
	} else {
		tmp := append(*this.dptr, []byte(s)...)
		this.dptr = &tmp
	}
	return this
}

func (this *_ByteArray) AppendByte(ch byte) ByteArray {
	tmp := append(*this.dptr, ch)
	this.dptr = &tmp
	return this
}

func (this *_ByteArray) AppendBytes(b []byte) ByteArray {
	tmp := append(*this.dptr, b...)
	this.dptr = &tmp
	return this
}

func (this *_ByteArray) At(i int) byte {
	if i <= len(*this.dptr) {
		return (*this.dptr)[i]
	}
	return 0
}

func (this *_ByteArray) Back() byte {
	if len(*this.dptr) > 0 {
		return (*this.dptr)[len(*this.dptr)-1]
	}
	return 0
}

// Removes n bytes from the end of the byte array.
func (this *_ByteArray) Chop(n int) {
	if n > len(*this.dptr) {
		tmp := (*this.dptr)[0:0]
		this.dptr = &tmp
		return
	}
	tmp := (*this.dptr)[0 : len(*this.dptr)-n]
	this.dptr = &tmp
}

func (this *_ByteArray) Chopped(n int) ByteArray {
	return NewByteArray((*this.dptr)[0 : len(*this.dptr)-n])
}

func (this *_ByteArray) Clear() {
	tmp := (*this.dptr)[0:0]
	this.dptr = &tmp
}

func (this *_ByteArray) ContainsOther(other ByteArray) bool {
	return bytes.Contains(*this.dptr, *other.dptr)
}

func (this *_ByteArray) ContainsString(str string) bool {
	return bytes.Contains(*this.dptr, []byte(str))
}

func (this *_ByteArray) ContainsByte(ch byte) bool {
	return bytes.Contains(*this.dptr, []byte{ch})
}

func (this *_ByteArray) CountOther(ba ByteArray) int {
	return bytes.Count(*this.dptr, *ba.dptr)
}

func (this *_ByteArray) CountString(str string) int {
	return bytes.Count(*this.dptr, []byte(str))
}

func (this *_ByteArray) CountByte(ch byte) int {
	return bytes.Count(*this.dptr, []byte{ch})
}

func (this *_ByteArray) Count() int   { return len(*this.dptr) }
func (this *_ByteArray) Data() []byte { return *this.dptr }

func (this *_ByteArray) EndsWithOther(ba ByteArray) bool {
	return bytes.HasSuffix(*this.dptr, *ba.dptr)
}

func (this *_ByteArray) EndsWithByte(ch byte) bool {
	return bytes.HasSuffix(*this.dptr, []byte{ch})
}

func (this *_ByteArray) EndsWithString(str string) bool {
	return bytes.HasSuffix(*this.dptr, []byte(str))
}

func (this *_ByteArray) Fill(ch byte, size ...int) ByteArray {
	if len(size) == 0 {
		for i := 0; i < len(*this.dptr); i++ {
			(*this.dptr)[i] = ch
		}
	} else {
		if size[0] <= len(*this.dptr) {
			tmp := (*this.dptr)[:size[0]]
			this.dptr = &tmp
			this.Fill(ch)
		} else {
			tmp := bytes.Repeat([]byte{ch}, size[0])
			this.dptr = &tmp
		}
	}
	return this
}

// Calling this function on an empty byte array constitutes undefined behavior.
func (this *_ByteArray) Front() byte { return (*this.dptr)[0] }

/*
func (this *_ByteArray) Mid(pos int, maxn ...int) ByteArray {
	return nil
}
*/
func (this *_ByteArray) IndexOfOther(ba ByteArray, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.Index((*this.dptr)[lowpos:], *ba.dptr)
}
func (this *_ByteArray) IndexOfString(str string, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.Index((*this.dptr)[lowpos:], []byte(str))
}
func (this *_ByteArray) IndexOfByte(ch byte, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.IndexByte((*this.dptr)[lowpos:], ch)
}
func (this *_ByteArray) IndexAny(chars string, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.IndexAny((*this.dptr)[lowpos:], chars)
}

func (this *_ByteArray) InsertOther(i int, ba ByteArray) ByteArray {
	tmp := append((*this.dptr)[:i], append(BytesDup(*ba.dptr), (*this.dptr)[i:]...)...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) InsertNChars(i int, count int, ch byte) ByteArray {
	ba := bytes.Repeat([]byte{ch}, count)
	tmp := append((*this.dptr)[:i], append(ba, (*this.dptr)[i:]...)...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) InsertString(i int, str string, length ...int) ByteArray {
	if len(length) == 0 {
		tmp := append((*this.dptr)[:i], append([]byte(str), (*this.dptr)[i:]...)...)
		this.dptr = &tmp
	} else {
		this.InsertString(i, StrSuf(str, length[0]))
	}
	return this
}
func (this *_ByteArray) InsertByte(i int, ch byte) ByteArray {
	tmp := append((*this.dptr)[:i], append([]byte{ch}, (*this.dptr)[i:]...)...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) IsEmpty() bool { return len(*this.dptr) == 0 }
func (this *_ByteArray) IsNull() bool  { return len(*this.dptr) == 0 }

func (this *_ByteArray) LastIndexOfOther(ba ByteArray, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.LastIndex((*this.dptr)[lowpos:], *ba.dptr)
}
func (this *_ByteArray) LastIndexOfString(str string, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.LastIndex((*this.dptr)[lowpos:], []byte(str))
}
func (this *_ByteArray) LastIndexOfByte(ch byte, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.LastIndexByte((*this.dptr)[lowpos:], ch)
}
func (this *_ByteArray) LastIndexOfAny(chars string, from ...int) int {
	lowpos := IfElseInt(len(from) == 0, 0, from[0])
	return bytes.LastIndexAny((*this.dptr)[lowpos:], chars)
}
func (this *_ByteArray) Left(maxlen int) ByteArray {
	highpos := IfElseInt(maxlen > len(*this.dptr), len(*this.dptr), maxlen)
	return NewByteArray((*this.dptr)[:highpos])
}
func (this *_ByteArray) LeftJustified(width int, truncate bool, fill ...byte) ByteArray {
	padsz := IfElseInt(len(*this.dptr) >= width, 0, width-len(*this.dptr))
	fillch := IfElse(len(fill) == 0, ' ', fill[0]).(byte)
	padarr := bytes.Repeat([]byte{fillch}, padsz)
	tmp := append(*this.dptr, padarr...)
	tmp = IfElse(truncate, tmp[:width], tmp).([]byte)
	return &_ByteArray{&tmp}
}
func (this *_ByteArray) Length() int { return len(*this.dptr) }
func (this *_ByteArray) Mid(pos int, maxlen ...int) ByteArray {
	if len(maxlen) == 0 {
		tmp := (*this.dptr)[pos:]
		return &_ByteArray{&tmp}
	} else {
		if pos+maxlen[0] > len(*this.dptr) {
			tmp := (*this.dptr)[pos:]
			return &_ByteArray{&tmp}
		} else {
			tmp := (*this.dptr)[pos : pos+maxlen[0]]
			return &_ByteArray{&tmp}
		}
	}
}

func (this *_ByteArray) PrependOther(ba ByteArray) ByteArray {
	tmp := append(BytesDup(*ba.dptr), *this.dptr...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) PrependNChars(count int, ch byte) ByteArray {
	barr := bytes.Repeat([]byte{ch}, count)
	tmp := append(barr, *this.dptr...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) PrependString(str string, maxlen ...int) ByteArray {
	tmp := append([]byte(str), *this.dptr...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) PrependData(data []byte, maxlen ...int) ByteArray {
	tmp := append(BytesDup(data), *this.dptr...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) PrependByte(ch byte) ByteArray {
	tmp := append([]byte{ch}, *this.dptr...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) Remove(pos, maxlen int) ByteArray {
	tmp := append((*this.dptr)[0:pos], (*this.dptr)[pos+maxlen:]...)
	this.dptr = &tmp
	return this
}
func (this *_ByteArray) Repeated(times int) ByteArray {
	tmp := bytes.Repeat(*this.dptr, times)
	return &_ByteArray{&tmp}
}

func (this *_ByteArray) ReplaceOther(pos, maxlen int, after ByteArray) ByteArray {

	return this
}

func (this *_ByteArray) Reverse() ByteArray {
	return nil
}

/*
QByteArray &	replace(int pos, int len, const QByteArray &after)
QByteArray &	replace(int pos, int len, const char *after, int alen)
QByteArray &	replace(int pos, int len, const char *after)
QByteArray &	replace(char before, const char *after)
QByteArray &	replace(char before, const QByteArray &after)
QByteArray &	replace(const char *before, const char *after)
QByteArray &	replace(const char *before, int bsize, const char *after, int asize)
QByteArray &	replace(const QByteArray &before, const QByteArray &after)
QByteArray &	replace(const QByteArray &before, const char *after)
QByteArray &	replace(const char *before, const QByteArray &after)
QByteArray &	replace(char before, char after)
QByteArray &	replace(const QString &before, const char *after)
QByteArray &	replace(char before, const QString &after)
QByteArray &	replace(const QString &before, const QByteArray &after)
void 	reserve(int size)
void 	resize(int size)
QByteArray 	right(int len) const
QByteArray 	rightJustified(int width, char fill = ' ', bool truncate = false) const
QByteArray &	setNum(int n, int base = 10)
QByteArray &	setNum(ushort n, int base = 10)
QByteArray &	setNum(short n, int base = 10)
QByteArray &	setNum(uint n, int base = 10)
QByteArray &	setNum(qlonglong n, int base = 10)
QByteArray &	setNum(qulonglong n, int base = 10)
QByteArray &	setNum(float n, char f = 'g', int prec = 6)
QByteArray &	setNum(double n, char f = 'g', int prec = 6)
QByteArray &	setRawData(const char *data, uint size)
void 	shrink_to_fit()
QByteArray 	simplified() const
int 	size() const
QList<QByteArray> 	split(char sep) const
void 	squeeze()
bool 	startsWith(const QByteArray &ba) const
bool 	startsWith(char ch) const
bool 	startsWith(const char *str) const
void 	swap(QByteArray &other)
QByteArray 	toBase64() const
QByteArray 	toBase64(QByteArray::Base64Options options) const
CFDataRef 	toCFData() const
double 	toDouble(bool *ok = nullptr) const
float 	toFloat(bool *ok = nullptr) const
QByteArray 	toHex() const
QByteArray 	toHex(char separator) const
int 	toInt(bool *ok = nullptr, int base = 10) const
long 	toLong(bool *ok = nullptr, int base = 10) const
qlonglong 	toLongLong(bool *ok = nullptr, int base = 10) const
QByteArray 	toLower() const
NSData *	toNSData() const
QByteArray 	toPercentEncoding(const QByteArray &exclude = QByteArray(), const QByteArray &include = QByteArray(), char percent = '%') const
CFDataRef 	toRawCFData() const
NSData *	toRawNSData() const
short 	toShort(bool *ok = nullptr, int base = 10) const
std::string 	toStdString() const
uint 	toUInt(bool *ok = nullptr, int base = 10) const
ulong 	toULong(bool *ok = nullptr, int base = 10) const
qulonglong 	toULongLong(bool *ok = nullptr, int base = 10) const
ushort 	toUShort(bool *ok = nullptr, int base = 10) const
QByteArray 	toUpper() const
QByteArray 	trimmed() const
void 	truncate(int pos)
const char *	operator const char *() const
const void *	operator const void *() const
bool 	operator!=(const QString &str) const
QByteArray &	operator+=(const QByteArray &ba)
QByteArray &	operator+=(const char *str)
QByteArray &	operator+=(char ch)
QByteArray &	operator+=(const QString &str)
bool 	operator<(const QString &str) const
bool 	operator<=(const QString &str) const
QByteArray &	operator=(const QByteArray &other)
QByteArray &	operator=(const char *str)
QByteArray &	operator=(QByteArray &&other)
bool 	operator==(const QString &str) const
bool 	operator>(const QString &str) const
bool 	operator>=(const QString &str) const
QByteRef 	operator[](int i)
char 	operator[](uint i) const
char 	operator[](int i) const
QByteRef 	operator[](uint i)
*/
/*
func ByteArrayFromBase64() ByteArray {

}
func ByteArrayFromHex() ByteArray {

}
func ByteArrayNumber() ByteArray {

}
*/
