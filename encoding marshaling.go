# json
package main
   
import (
    "fmt"
    "encoding/json"
)
   

type Human struct{
          
    Name string
    Age int
    Address string
}
   
func main() {
       
    
    human1 := Human{"Ankit", 23, "New Delhi"}
         
    human_enc, err := json.Marshal(human1)
       
    if err != nil {
           
        
        fmt.Println(err)
    }
       
    fmt.Println(string(human_enc))
    
    human2 := []Human{
        {Name: "vinodhini", Age: 22, Address: "New Delhi"},
        {Name: "enbam", Age: 27, Address: "Pune"},
        {Name: "deepa", Age: 21, Address: "Bangalore"},
    }
       
    
    human2_enc, err := json.Marshal(human2)
        
        if err != nil {
       
        
            fmt.Println(err)
        }
           
    
    fmt.Println()
        fmt.Println(string(human2_enc))
}
