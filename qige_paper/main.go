package main
//这个代码都可以再修改整合一下
import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	//filePath := "C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/Final_14.5/seismic_"
	//filePath := "C:/Users/JACOB/Desktop/21-10-17-datacode-qige/22-2-18/section3.3.2-backfill type2-baohe/"
	filePath := "E:/data/Wenchuan"
	ObjFile := "E:/data/EL"
	for i := 1;i<=1 ;i++{
		i_string := strconv.Itoa(i)
		for j:= 1;j<=1;j++{
			j_string := strconv.Itoa(j)
			//getAccx(filePath,i_string,j_string)
			getMoment(filePath,i_string,j_string,ObjFile)
			getSoilXdisp(filePath,i_string,j_string,ObjFile)
			//getXvel(filePath,i_string,j_string)
			//getTunnelXdisp(filePath,i_string,j_string)
			//getTunnelYdisp(filePath,i_string,j_string)
			getAxial(filePath,i_string,j_string,ObjFile)
			getShear(filePath,i_string,j_string,ObjFile)
			//getMaxDisp(filePath,i_string,j_string)
		}
	}
}

func getSoilXdisp(filePath,i_string string,j_string string,ObjFile string) {
	node := 2
	filepath :=filePath+"-0."+j_string+"g-xdis-ling.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
		return
	}
	lines := strings.Split(string(b),"\n")
	fmt.Println(len(lines),i_string,j_string)
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	resData := getDispData(lines,node)
	//写入文件夹。
	filetitle := "EL_0."+j_string+"_xdisp_ling"
	objFile := ObjFile+"_xdisp_ling_max.csv"
	headLine := []string{"element.no", "max_disp",  "abs_disp"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getDispData(lines []string, node int) [][]string{
	node2:= node
	linescount := len(lines)
	fmt.Println(linescount)
	var eachNodeCount int
	if linescount/node%1 != 0 {
		eachNodeCount = linescount/node + 1
	} else {
		eachNodeCount = linescount / node
	}
	fmt.Println(eachNodeCount)
	var ResDate [][]string
	var xD1 [][]float64
	for j := 0; j < node2; j++ {
		start := j*eachNodeCount + 4
		end := (j + 1) * eachNodeCount
		var lines2 []float64
		for i := start; i < end; i++ {
			//把这txt给读出来。
			data := strings.Replace(lines[i], " ", ",", -1)
			data2 := strings.Split(data, ",,,,,,,,,,,,")
			data3 := strings.Replace(data2[1], ",", "", -1)
			//fmt.Println(data3)
			data4 := strings.Replace(data3, "\r", "", -1)
			dataint, err := strconv.ParseFloat(data4, 64)
			if err != nil {
				log.Fatalf("can not convert, err is %+v", err)
			}
			lines2 = append(lines2, dataint)
		}
		xD1 = append(xD1,lines2)
	}
	for k:=0;k<1;k++ {
		var max float64 = 0
		var _disp float64
		for j:= 0;j<len(xD1[k]);j++{
			disp := xD1[k][j]-xD1[k+1][j]
			if math.Abs(disp)>max{
				max=math.Abs(disp)
				_disp = disp
			}
		}
		node = k+1
		node_string := strconv.Itoa(node)
		max_string := strconv.FormatFloat(max,'E',-1,64)
		disp_string := strconv.FormatFloat(_disp,'E',-1,64)
		ResDate = append(ResDate,[]string{
			node_string,disp_string,max_string,
		})
	}
	return ResDate
}
func getMaxData(xlines []string, ylines []string, node int) [][]string{
	node2:= node
	linescount := len(xlines)
	fmt.Println(linescount)
	var eachNodeCount int
	if linescount/node%1 != 0 {
		eachNodeCount = linescount/node + 1
	} else {
		eachNodeCount = linescount / node
	}
	fmt.Println(eachNodeCount)
	var ResDate [][]string
	var xData,yData [][]string
	//xData := make(map[int]floa64t64,len(xlines))
	//yData := make(map[int]float64,len(ylines))
	for j := 0; j < node2; j++ {
		start := j*eachNodeCount + 4
		end := (j + 1) * eachNodeCount
		var xdata,y2data []string
		for i := start; i < end; i++ {
			//把这txt给读出来。x的disp
			data := strings.Replace(xlines[i], " ", ",", -1)
			data2 := strings.Split(data, ",,,,,,,,,,,,")
			data3 := strings.Replace(data2[1], ",", "", -1)
			data4 := strings.Replace(data3, "\r", "", -1)
			ydata := strings.Replace(ylines[i], " ", ",", -1)
			ydata2 := strings.Split(ydata, ",,,,,,,,,,,,")
			ydata3 := strings.Replace(ydata2[1], ",", "", -1)
			ydata4 := strings.Replace(ydata3, "\r", "", -1)
			xdata = append(xdata,data4)
			y2data =append(y2data,ydata4)
		}
		xData = append(xData,xdata)
		yData = append(yData,y2data)
	}
	for i:= 0;i<16;i++{
		var max float64 = 0
		var x1_string,y1_string,x2_string,y2_string string
		for j:=0;j<len(xData[i]);j++{
			x1,_ := strconv.ParseFloat(xData[i][j],64)
			y1,_ := strconv.ParseFloat(yData[i][j],64)
			x2,_ := strconv.ParseFloat(xData[i+16][j],64)
			y2,_ := strconv.ParseFloat(yData[i+16][j],64)
			x := x1-x2
			y := y1-y2
			dis := math.Sqrt(x*x+y*y)
			if dis> max{
				max = dis
				x1_string =strconv.FormatFloat(x1,'E',-1,64)
				y1_string =strconv.FormatFloat(y1,'E',-1,64)
				x2_string =strconv.FormatFloat(x2,'E',-1,64)
				y2_string =strconv.FormatFloat(y2,'E',-1,64)
			}
		}
		node = i + 1
		node_string := strconv.Itoa(node)//这个nodenumber跟时程里面的不一样。
		max_string := strconv.FormatFloat(max, 'E', -1, 64)
		//min_string := strconv.FormatFloat(min, 'E', -1, 64)
		//absMax_string := strconv.FormatFloat(absMax, 'E', -1, 64)
		fmt.Printf("node:%d; max:%f\n ", node, max)
		//fmt.Printf("node:%d; min:%f\n ", node, min)
		ResDate = append(ResDate, []string{node_string, max_string,x1_string,y1_string,x2_string,y2_string})
	}
	return ResDate
}



func getShear(filePath,i_string string, j_string string,ObjFile string) {
	node := 8
	filepath :=filePath+"-0."+j_string+"g-shear.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
		return
	}
	lines := strings.Split(string(b),"\n")
	fmt.Println(len(lines),i_string,j_string)
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	resData := getData(lines,node)
	//写入文件夹。
	filetitle := "EL_0."+j_string+"_shear"
	objFile :=ObjFile+ "_0.1g_shear_max.csv.csv"
	headLine := []string{"element.no", "max", "min", "Max_shear"}
	writeFile(objFile,filetitle,resData,headLine)
}




func getAxial(filePath,i_string string,j_string string,ObjFile string) {
	node := 8
	filepath :=filePath+"-0."+j_string+"g-axial.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
		return
	}
	lines := strings.Split(string(b),"\n")
	fmt.Println(len(lines),i_string,j_string)
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	resData := getData(lines,node)
	//写入文件夹。
	filetitle := "EL_0."+j_string+"_axial"
	objFile := ObjFile+"_0.1g_axial_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_axial"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getMoment(filePath,i_string string,j_string string,ObjFile string) {
	node := 8
	filepath :=filePath+"-0."+j_string+"g-moment1.txt"
	b,e := ioutil.ReadFile(filepath)
	if e !=nil{
		fmt.Println("read file error")
		return
	}
	lines := strings.Split(string(b),"\n")
	fmt.Println(len(lines),i_string,j_string)
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	resData := getData(lines,node)
	//写入文件夹。
	filetitle := "EL_0."+j_string+"_moment"
	objFile := ObjFile+"_0.1g_moment1_max.csv.csv"
	headLine := []string{"element.no", "max", "min", "Max_moment"}
	writeFile(objFile,filetitle,resData,headLine)
}

func writeFile(file,filetitle string,resData [][]string,headline []string) {
	nfs, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot create file,err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	var title []string
	title = append(title, filetitle)
	err = w.Write(title)
	if err != nil {
		log.Fatalf("cannot write title：%v", err)
	}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	//一次写入多行
	w.WriteAll(resData)
}

func getData(lines []string,node int) [][]string {
	node2:= node
	linescount := len(lines)
	fmt.Println(linescount)
	var eachNodeCount int
	if linescount/node%1 != 0 {
		eachNodeCount = linescount/node + 1
	} else {
		eachNodeCount = linescount / node
	}
	fmt.Println(eachNodeCount)
	var ResDate [][]string
	for j := 0; j < node2; j++ {
		start := j*eachNodeCount + 4
		end := (j + 1) * eachNodeCount
		var lines2 []float64
		for i := start; i < end; i++ {
			//把这txt给读出来。
			data := strings.Replace(lines[i], " ", ",", -1)
			data2 := strings.Split(data, ",,,,,,,,,,,,")
			data3 := strings.Replace(data2[1], ",", "", -1)
			//fmt.Println(data3)
			data4 := strings.Replace(data3, "\r", "", -1)
			dataint, err := strconv.ParseFloat(data4, 64)
			if err != nil {
				log.Fatalf("can not convert, err is %+v", err)
			}
			lines2 = append(lines2, dataint)
		}
		sort.Float64s(lines2)
		node = j + 1
		max := lines2[len(lines2)-1]
		min := lines2[0]
		abs := math.Abs(min)
		var absMax float64
		if max >= abs {
			absMax = max
		} else {
			absMax = min
		}
		node_string := strconv.Itoa(node)//这个nodenumber跟时程里面的不一样。
		max_string := strconv.FormatFloat(max, 'E', -1, 64)
		min_string := strconv.FormatFloat(min, 'E', -1, 64)
		absMax_string := strconv.FormatFloat(absMax, 'E', -1, 64)
		fmt.Printf("node:%d; max:%f\n ", node, max)
		fmt.Printf("node:%d; min:%f\n ", node, min)
		ResDate = append(ResDate, []string{node_string, max_string, min_string, absMax_string})
	}
	return ResDate
}
