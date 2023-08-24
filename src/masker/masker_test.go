package masker

import (
	"fmt"
	"testing"

	"github.com/showa-93/go-mask"
)

func TestMask(t *testing.T) {
	{
		maskValue, _ := mask.String(mask.MaskTypeFixed, "Hello World!!")
		fmt.Println(maskValue)
	}
	{
		value := struct {
			Title string   `mask:"filled"`
			Casts []string `mask:"fixed"`
		}{
			Title: "Catch Me If You Can",
			Casts: []string{
				"Thomas Jeffrey \"Tom\" Hanks",
				"Leonardo Wilhelm DiCaprio",
			},
		}
		//masker := mask.NewMasker()
		//masker.SetMaskChar("-")
		//masker.RegisterMaskStringFunc(mask.MaskTypeFilled, masker.MaskFilledString)
		//masker.RegisterMaskStringFunc(mask.MaskTypeFixed, masker.MaskFixedString)

		maskValue, _ := mask.Mask(value)
		//maskValue2, _ := masker.Mask(value)

		fmt.Printf("%+v\n", value)
		fmt.Printf("%+v\n", maskValue)
		//fmt.Printf("%+v\n", maskValue2)

	}
}
