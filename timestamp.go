now := time.Now()      
sec := now.Unix()      
nsec := now.UnixNano() 

fmt.Println(now) 
fmt.Println(sec)
fmt.Println(nsec) 
