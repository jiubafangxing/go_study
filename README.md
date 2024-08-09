Golang：从0到1的学习笔记

疑惑： for range 模式的遍历，接受参数时，有时直接是一个，有时两个，如何得知用几个参数接受

```
 func add(args []reflect.Value)(results []reflect.Value){
        if len(args) == 0{
                return nil
        }
        var ret  reflect.Value
        switch args[0].Kind(){
        case reflect.Int:
                n := 0
                for _, v := range args{
                        n+=int(v.Int()) 
                }
                ret = reflect.ValueOf(n)

        case reflect.String:
                ss := make ([]string, 0, len(args))
                for _,v := range args{
                        ss = append(ss, v.String())
                }
                ret = reflect.ValueOf(strings.Join(ss, ""))

        }
        results =  append(results, ret)
        return results
}

```
以上片段中，为什么 n+=int(v.Int()) 需要这么一转呢
