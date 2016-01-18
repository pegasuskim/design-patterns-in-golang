package db
import "fmt"

type repository struct {
   items map[string]string
}

func (r *repository) Get(key string) (string, error) {
   item, ok := r.items[key]
   if !ok {
      return "", fmt.Errorf("The '%s' is not presented", key)
   }
   return item, nil
}


var (
   r    *repository
   once sync.Once
)

func Repository() *repository {
   if r == nil {
      once.Do(func() {
         r = &repository{
            items: make(map[string]string),
         }
      })
   }
   return r
}
