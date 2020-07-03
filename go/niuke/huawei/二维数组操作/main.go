package main

import(
    "fmt"
)

func main(){
    var n, m int
    
    for {
        _,err:=fmt.Scan(&n, &m)
        if err!=nil{
            return 
		}
		str:=""
		if n<=9 && m<=9{
			str+="0\n"
		}else{
			str+="-1\n"
		}
        
        var x, y, a,b int
        fmt.Scan(&x, &y, &a,&b)
        if x>n-1 || y>m-1 || a>n-1 || b>m-1{
            str+="-1\n"
		}else{
			str+="0\n"
		}
        
        fmt.Scan(&x)
        if x < n{
            str+="0\n"
        }else{
            str+="-1\n"
        }
        
        fmt.Scan(&x)
        if x < m{
            str+="0\n"
        }else{
            str+="-1\n"
        }
        
        fmt.Scan(&x,&y)
        if x < n && y<m{
            str+="0\n"
        }else{
            str+="-1\n"
        }
        
        fmt.Printf("%s", str)
    }
}