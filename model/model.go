package model

type Effectiveness struct {
    Total int
    Correct int
    Effectiveness float64
}

type Auth struct  {
    Auth string `schema:"auth"`
}

type Comment struct {
    Name string `schema:"name"`
    Comment string `schema:"msg"`
}
