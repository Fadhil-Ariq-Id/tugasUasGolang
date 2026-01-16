packagae main

import {
"fmt"
"time"
}


func main(){
//  add gabungang  graf dari data  statis
g := BuildGraph()
source :=SourceCode()

// run djisktra  dan measeru time
start := time.now()
dist, prev := Dijkstra(g, source)
elapsed  := time.Since.(start)
}
