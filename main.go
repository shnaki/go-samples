package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		{"p01", "free", "0.15", "511", "1", "0", "0", "6/3", "0", ""},
		{"p02", "free", "0.00", "511", "1", "0", "0", "0/0", "0", ""},
		{"p03", "free", "0.00", "511", "1", "0", "0", "0/0", "0", ""},
		{"p04", "free", "0.00", "511", "1", "0", "0", "0/0", "0", ""},
		{"p05", "free", "0.00", "511", "1", "0", "0", "0/0", "0", ""},
		{"p06", "free", "0.00", "511", "1", "0", "0", "0/0", "0", ""},
		{"p07", "excl", "0.73", "511", "1", "0", "0", "1/1", "0", "219"},
		{"p08", "excl", "0.94", "511", "1", "0", "0", "1/1", "0", "217"},
		{"p09", "excl", "0.99", "511", "1", "0", "0", "1/1", "0", "217"},
	}

	// Original output is below.
	//node state  load    pmem ncpu   mem   resi usrs jobs   jobids
	// p01  free  0.15     511   1      0      0  6/3    0
	// p02  free  0.00     511   1      0      0  0/0    0
	// p03  free  0.00     511   1      0      0  0/0    0
	// p04  free  0.00     511   1      0      0  0/0    0
	// p05  free  0.00     511   1      0      0  0/0    0
	// p06  free  0.00     511   1      0      0  0/0    0
	// p07  excl  0.73     511   1      0      0  1/1    1   219
	// p08  excl  0.94     511   1      0      0  1/1    1   217
	// p09  excl  0.99     511   1      0      0  1/1    1   217

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderLine(false)
	table.SetHeader([]string{"node", "state", "load", "pmem", "ncpu", "mem", "resi", "usrs", "jobs", "jobids"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)
	//table.SetCenterSeparator("")
	//table.SetColumnSeparator("")
	//table.SetRowSeparator("")
	table.SetNoWhiteSpace(true)
	table.SetTablePadding(" ")
	table.AppendBulk(data)
	table.Render()
}
