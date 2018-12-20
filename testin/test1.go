package testin

var Day1Part1Tests = []string{
`+1
+1
+1`,

`+1
+1
-2`,

`-1
-2
-3`,
}
var Day1Part1Expects = []string{`3`, `0`, `-6`}

var Day1Part2Tests = []string{
`+1
-1`,

`+3
+3
+4
-2
-4`,

`-6
+3
+8
+5
-6`,

`+7
+7
-2
-7
-4`,
}

var Day1Part2Expects = []string{`0`, `10`, `5`, `14`}