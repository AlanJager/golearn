package main

import ( 
    "fmt"
    "errors"
)

func main(){
    var frenchHello string
    
    no, yes, maybe := "no", "yes", "maybe"
    japaneseHello := "Konichiwa"
    frenchHello = "Bonjour"
    fmt.Printf("no is %s, yes is %s, maybe is %s, japaneseHello is %s, frenchHello is %s", no, yes, maybe, japaneseHello, frenchHello)
     
    s := "hello"
    c := []byte(s)
    c[0] = 'c'
    s2 := string(c)
    fmt.Printf("%s\n", s2)

    s = "hello,"
    m := " world"
    a := s + m
    fmt.Printf("%s\n", a)

    s = "hello"
    s = "c" + s[:1]
    fmt.Printf("%s\n", s)

    m = `hello
           world`
    fmt.Printf("%s\n", m)

    err := errors.New("emit macho dwarf: elf header corrupted")
    if err != nil {
        fmt.Print(err)
    }

} 
