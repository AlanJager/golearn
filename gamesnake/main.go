package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

import "C" 

type loct struct {
    i, j int
}

var (
    area = [20][20]byte{} 
    food bool             
    lead byte            
    head loct            
    tail loct             
    size int              
)

func place() loct {
    k := rand.Int() % 400
    return loct{k / 20, k % 20}
}

func draw(p loct, c byte) {
    C.gotoxy(C.int(p.i*2+4), C.int(p.j+2))
    fmt.Fprintf(os.Stderr, "%c", c)
}

func init() {

    head, tail = loct{4, 4}, loct{4, 4}
    lead, size = 'R', 1
    area[4][4] = 'H'
    rand.Seed(int64(time.Now().Unix()))

    fmt.Fprintln(os.Stderr,
        `
  #-----------------------------------------#
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |         *                               |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  |                                         |
  #-----------------------------------------#
`)

    go func() {
        for { 
            switch byte(C.direct()) {
            case 72:
                lead = 'U'
            case 75:
                lead = 'L'
            case 77:
                lead = 'R'
            case 80:
                lead = 'D'
            case 32:
                lead = 'P'
            }
        }
    }()
}

func main() {

    for {
        time.Sleep(time.Millisecond * 400)

        if lead == 'P' {
            continue
        }

        if !food {
            give := place()
            if area[give.i][give.j] == 0 { 
                area[give.i][give.j] = 'F'
                draw(give, '$') 
                food = true
            }
        }

        area[head.i][head.j] = lead

        switch lead {
        case 'U':
            head.j--
        case 'L':
            head.i--
        case 'R':
            head.i++
        case 'D':
            head.j++
        }

        if head.i < 0 || head.i >= 20 || head.j < 0 || head.j >= 20 {
            C.gotoxy(0, 23) 
            break 
        }

        eat := area[head.i][head.j]

        if eat == 'F' {
            food = false

            size++
        } else if eat == 0 { 

            draw(tail, ' ') 

            dir := area[tail.i][tail.j]

            area[tail.i][tail.j] = 0

            switch dir {
            case 'U':
                tail.j--
            case 'L':
                tail.i--
            case 'R':
                tail.i++
            case 'D':
                tail.j++
            }
        } else { 
            C.gotoxy(0, 23)
            break
        }
        draw(head, '*') 
    }

    switch {
    case size < 22:
        fmt.Fprintf(os.Stderr, "Faild! You've eaten %d $\\n", size-1)
    case size < 42:
        fmt.Fprintf(os.Stderr, "Try your best! You've eaten %d $\\n", size-1)
    default:
        fmt.Fprintf(os.Stderr, "Congratulations! You've eaten %d $\\n", size-1)
    }
}
//该片段来自于http://outofmemory.cn
