# Book-Like-Pdf-Maker
- 両面印刷非対応のプリンタで両面印刷をするためにPDFファイルを分割するためのソフトウェアです。
- おそらくこんな面倒なことをしなくてもできるソフトウェアがあります・・・。

# Installation
- go get github.com/cs3238-tsuzu/book-like-pdf-maker

# Usage
- book-like-pdf-maker "path to file"
- OKと出力されれば成功です。ファイルがあったフォルダに\_odd.pdfと\_even.pdfで終わる2つのファイルが生成されます。
- \_odd.pdfで終わるファイルを開き、印刷します。
- 印刷された紙を表裏逆にし、再びセットします。
- \_even.pdfで終わるファイルを開き印刷します。
- 印刷された紙をステープラ等で端を止めれば完成です。

# License
- Under the MIT License
- Copyright (c) 2017 Tsuzu
