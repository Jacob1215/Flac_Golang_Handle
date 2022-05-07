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
	//filePath := "C:/Users/JACOB/Desktop/21-10-17-datacode-qige/22-2-18/backfill grouting-type2-2-16/"
	//filePath := "G:/res_201-300_2/seismic_"
	filePath := "C:/Users/JACOB/Desktop/seismic_wave/4-8_recalc/seismic_"
	//filePath := "C:/Users/JACOB/Desktop/sei
	//smic_wave/static_load/max_free_field_displacement/seismic_"
	//ObjFile := "C:/Users/JACOB/Desktop/seismic_wave/static_load/max_moment_env"
	ObjFile := "C:/Users/JACOB/Desktop/seismic_wave/4-8_recalc"
	for i := 1;i<= 1;i++{
		i_string := strconv.Itoa(i)
		for j:= 10;j<=10;j++{
			j_string := strconv.Itoa(j)
			//getAccx(filePath,i_string,j_string,ObjFile)
			//getMoment(filePath,i_string,j_string,ObjFile)
			//getSoilXdisp(filePath,i_string,j_string,ObjFile)
			//getXvel(filePath,i_string,j_string,ObjFile)
			getTunnelXdisp(filePath,i_string,j_string,ObjFile)
			getTunnelYdisp(filePath,i_string,j_string,ObjFile)
			//getAxial(filePath,i_string,j_string,ObjFile)
			//getShear(filePath,i_string,j_string,ObjFile)
			//maxStep := getMaxDisp(filePath,i_string,j_string,ObjFile)
			//getMaxEvp(filePath,i_string,j_string,ObjFile,maxStep)
		}
	}
}

func getMaxEvp(path string, i_string string, j_string string, file string, timeStep int) {
	node := 32
	filepath :=path+i_string+"_0."+j_string+"_moment.txt"
	x,_:= ioutil.ReadFile(filepath)
	lines := strings.Split(string(x),"\n")
	//fmt.Println(len(lines),i_string,j_string)
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
	var timeStep2,timeStep3 int
	for j := 0; j < node2; j++ {
		start := j*eachNodeCount + 4
		end := (j + 1) * eachNodeCount
		var lines2 []float64
		for i := start; i < end; i++ {
			//把这txt给读出来。
			data := strings.Replace(lines[i], " ", ",", -1)
			//fmt.Println(data)
			data2 := strings.Split(data, ",,,,,,,,,,")
			//fmt.Println(data2[0],data2[1],data2[2])
			data3 := strings.Replace(data2[len(data2)-1], ",", "", -1)
			//fmt.Println(data3)
			//fmt.Println(data3)
			data4 := strings.Replace(data3, "\r", "", -1)
			dataint, err := strconv.ParseFloat(data4, 64)
			if err != nil {
				log.Fatalf("can not convert, err is %+v", err)
			}
			lines2 = append(lines2, dataint)
		}

		var max,min float64 = 0,0
		for i,v :=range lines2{
			if v > max {
				max = v
				timeStep2 = i
			}
			if v < min {
				min = v
				timeStep3 = i
			}
		}
		abs := math.Abs(min)
		var absMax float64
		if max >= abs {
			absMax = max
			timeStep2 = timeStep2
		} else {
			absMax = min
			timeStep2 = timeStep3
		}
		node_string := strconv.Itoa(j)//这个nodenumber跟时程里面的不一样。
		max_string := strconv.FormatFloat(max, 'E', -1, 64)
		min_string := strconv.FormatFloat(min, 'E', -1, 64)
		absMax_string := strconv.FormatFloat(absMax, 'E', -1, 64)
		maxMoment_step := strconv.FormatFloat(lines2[timeStep],'E',-1,64)
		maxMoment_step2 := strconv.FormatFloat(lines2[timeStep2],'E',-1,64)
		fmt.Printf("node:%d; max:%f\n ", node, max)
		fmt.Printf("node:%d; min:%f\n ", node, min)
		ResDate = append(ResDate, []string{node_string, max_string, min_string, absMax_string,maxMoment_step,maxMoment_step2 })
	}
	//写入文件夹。
	filetitle := "seismic_"+i_string+"_0."+j_string+"_moment"
	objFile := file+"/seismic"+i_string+"_moment_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_moment","Max_Moment_env/D","maxMoment_Env/M"}
	writeFile(objFile,filetitle,ResDate,headLine)
}

func getMaxDisp(filePath string, i_string string, j_string string,ObjFile string) int {
	node := 32
	filepath :=filePath+i_string+"_0."+j_string+"_tunnel_xdisp.txt"
	x,_:= ioutil.ReadFile(filepath)
	xlines := strings.Split(string(x),"\n")
	fmt.Println(len(xlines),i_string,j_string)
	filepathy :=filePath+i_string+"_0."+j_string+"_tunnel_ydisp.txt"
	y,_ := ioutil.ReadFile(filepathy)
	ylines := strings.Split(string(y),"\n")
	fmt.Println(len(ylines),i_string,j_string)
	resData,timeStep :=getMaxData(xlines,ylines,node)
	//计算node次//这样一个文件夹就搞定了。
	//获取数据集
	//写入文件夹。
	filetitle := "seismic_"+i_string+"_0."+j_string+"_tunnel"
	objFile := ObjFile+"/seismic"+i_string+"_tunnel_maxdisp_max.csv"
	headLine := []string{"element.no", "max", "x1_disp", "y1_disp","x2_","y2_"}
	writeFile(objFile,filetitle,resData,headLine)
	return timeStep
}

func getMaxData(xlines []string, ylines []string, node int) ([][]string,int){
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
			data2 := strings.Split(data, ",,,,,,,,,,")
			data3 := strings.Replace(data2[len(data2)-1], ",", "", -1)
			data4 := strings.Replace(data3, "\r", "", -1)
			ydata := strings.Replace(ylines[i], " ", ",", -1)
			ydata2 := strings.Split(ydata, ",,,,,,,,,,")
			ydata3 := strings.Replace(ydata2[len(data2)-1], ",", "", -1)
			ydata4 := strings.Replace(ydata3, "\r", "", -1)
			xdata = append(xdata,data4)
			y2data =append(y2data,ydata4)
		}
		xData = append(xData,xdata)
		yData = append(yData,y2data)
	}
	var timeStep int
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
				timeStep = j
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
	return ResDate,timeStep
}

func getTunnelYdisp(filePath,i_string string, j_string string,ObjFile string) {
	node := 32
	filepath :=filePath+i_string+"_0."+j_string+"_tunnel_ydisp.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_tunnel_ydisp"
	objFile := ObjFile+"/seismic"+i_string+"_tunnel_ydisp_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_axial"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getTunnelXdisp(filePath,i_string string, j_string string,ObjFile string) {
	node := 32
	filepath :=filePath+i_string+"_0."+j_string+"_tunnel_xdisp.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_tunnel_xdisp"
	objFile := ObjFile+"/seismic"+i_string+"_tunnel_xdisp_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_axial"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getShear(filePath,i_string string, j_string string,ObjFile string) {
	node := 32
	filepath :=filePath+i_string+"_0."+j_string+"_shear.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_shear"
	objFile := ObjFile+"/shear/seismic"+i_string+"_shear_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_shear"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getXvel(filePath,i_string string,j_string string,ObjFile string) {
	node := 15
	filepath :=filePath+i_string+"_0."+j_string+"_xvel.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_xvel"
	objFile := ObjFile+"/xvel/seismic"+i_string+"_xvel_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_axial"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getSoilXdisp(filePath,i_string string,j_string string,ObjFile string) {
	node := 15
	filepath :=filePath+i_string+"_0."+j_string+"_soil_xdisp.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_soil_xdisp"
	objFile := ObjFile+"/seismic"+i_string+"_soil_xdisp_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_axial"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getAxial(filePath,i_string string,j_string string,ObjFile string) {
	node := 32
	filepath :=filePath+i_string+"_0."+j_string+"_axial.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_axial"
	objFile := ObjFile+"/axial/seismic"+i_string+"_axial_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_axial"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getMoment(filePath,i_string string,j_string string,ObjFile string) {
	node := 32
	filepath :=filePath+i_string+"_0."+j_string+"_moment.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_moment"
	objFile := ObjFile+"/moment/seismic"+i_string+"_moment_max.csv"
	headLine := []string{"element.no", "max", "min", "Max_moment"}
	writeFile(objFile,filetitle,resData,headLine)
}

func getAccx(filePath,i_string string,j_string string,ObjFile string) {
	node := 15 //good
	filepath :=filePath+i_string+"_0."+j_string+"_accx.txt"
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
	filetitle := "seismic_"+i_string+"_0."+j_string+"_accx"
	objFile := ObjFile+"/seismic"+i_string+"_accx_max.csv"
	headLine := []string{"node.no", "max", "min", "Max_accx"}
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
			//fmt.Println(data)
			data2 := strings.Split(data, ",,,,,,,,,,")
			//fmt.Println(data2[0],data2[1],data2[2])
			data3 := strings.Replace(data2[len(data2)-1], ",", "", -1)
			//fmt.Println(data3)
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
