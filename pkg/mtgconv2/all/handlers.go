package all

import (
	_ "mtgconv/pkg/mtgconv2/sources/moxfield"
	_ "mtgconv/pkg/mtgconv2/sources/archidekt"
	_ "mtgconv/pkg/mtgconv2/sources/shinycsv"
	_ "mtgconv/pkg/mtgconv2/sources/txt-moxfield"

	_ "mtgconv/pkg/mtgconv2/outputs/dck"
	_ "mtgconv/pkg/mtgconv2/outputs/txt"
	_ "mtgconv/pkg/mtgconv2/outputs/json"
	_ "mtgconv/pkg/mtgconv2/outputs/moxfield-collection"
)

// register all the mtgconv input and output handlers here
// so that you can just import this single package and have it init all of them at once