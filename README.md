# ruth
Use of try, catch, finally in Golang using the same syntax as java / python / php

<p align="center">
<img src="https://user-images.githubusercontent.com/39353278/72227565-313f4380-35af-11ea-8fc6-c33a5945b867.png" alt="ruth" width="300">
</p>

## Example


```go

  
	var array []string

	ruth.Try(func() {

		fmt.Print(array[0])

	}).Catch(func(e ruth.Exception) {

		fmt.Println("index error")

	}).Finally(func() {
		
		fmt.Println("finally")
		
	})


```
## Response
```
index error
finally
```



## Author

>Tolga Özen

>mtolgaozen@gmail.com


## License

MIT License

Copyright (c) 2020 Tolga Ozen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.