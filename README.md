# genrsa

RSA private and public keys generator written in Go


## 📝 Table of Contents

- [Build](#build)
- [Usage](#usage)
- [Example](#example)
- [Built Using](#built_using)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## Build <a name = "build"></a>

```bash
make build
```

## Usage <a name = "usage"></a>

genrsa [--bits number] [--output filename]
&nbsp;&nbsp;options:
&nbsp;&nbsp;&nbsp;&nbsp;--bits: optional number of bits. default 4096
&nbsp;&nbsp;&nbsp;&nbsp;--output: optional output file name. default "rsa"

Private key will be written to output file name with extension .pem
Public key will be written to output file name with extension .pub

## Example <a name ="example"></a>

Generating a RSA keys with bit size of 4096 to files my-key.pem and my-key.pub

```bash
./genrsa --bits 4069 --output my-key
```
## Built Using <a name = "built_using"></a>

- [Go](https://go.dev/) - Server Framework

## Authors <a name = "authors"></a>

- [@PaulShpilsher](https://github.com/PaulShpilsher) - Software engineer and digital nomad

## Acknowledgements <a name = "acknowledgement"></a>

- Inspiration ["Simple Made Easy" - Rich Hickey (2011)](https://www.youtube.com/watch?v=SxdOUGdseq4)
