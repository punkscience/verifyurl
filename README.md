# VerifyURL
A quick golang module/function to verify that a given URL is well formed and ensure that there's something valid on the other end.
## Installation
go get github.com/punkscience/verifyurl
## Usage
```
  package main
  
  import (
    "github.com/punkscience/verifyurl"
    "log"
  )
  
  func main() {
    myurl := "https://google.ca"
    
    err := verifyurl.Verify(myurl)
    if err != nil {
      log.Println( "Something has gone horribly wrong with Google.", err )
    }
  }
```
## Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D
## History
I needed it. I wrote it. I thought others might be able to use it so I made it a module.
## Credits
Created initially by Darryl Wright
## License
Free and open. Do whatever you wish. 
