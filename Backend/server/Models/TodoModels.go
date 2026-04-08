package Models
type Todo struct {
 Id      uint   `json:"id"`
 Title    string `json:"title"`
 Completed   bool `json:"comopleted"`
 Created_at   string `json:"created_at"`
 Updated_at string `json:"updated_at"`
}
func (b *Todo) TableName() string {
 return "todo"
}