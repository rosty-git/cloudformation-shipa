package resource

import "testing"

func Test_convertModel(t *testing.T) {
    name := "test"
    m := convertModel(&Model{
        Name: &name,
    })

    if m.Name != name {
        t.Error("name not expected")
    }
}
