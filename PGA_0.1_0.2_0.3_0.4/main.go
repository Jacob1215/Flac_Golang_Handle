package main

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"strconv"
)

func main()  {
	//get_waves_1st()
	get_waves_2ec()

}

func get_waves_2ec() {
	filepath3 := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/vel_History/I_PGD_PGV.xlsx"
	xlsx2, err := excelize.OpenFile(filepath3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//调数据的地方
	fileNO := 214
	end := 236
	fileNo_string :=strconv.Itoa(fileNO)

	//这里是调幅系数-第1次
	rows2 := xlsx2.GetRows("I_PGD_PGV" )
	Filename := fileNo_string+"PGa_real"
	t := 1
	TiaofuPath :="C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/all_seismic_tiaofu/"+Filename+".xlsx"
	xlsxTiaofu, err := excelize.OpenFile(TiaofuPath)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	rowsTiaofu := xlsxTiaofu.GetRows(Filename)
	//第二次计算调幅
	Filename3 := fileNo_string+"PGa_real_2"
	TiaofuPath3 :="C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/all_seismic_tiaofu/"+Filename3+".xlsx"
	xlsxTiaofu3, err := excelize.OpenFile(TiaofuPath3)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	rowsTiaofu3 := xlsxTiaofu3.GetRows(Filename3)
	//第三次计算调幅
	Filename4 := fileNo_string+"PGa_real_3"
	TiaofuPath4 :="C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/all_seismic_tiaofu/"+Filename4+".xlsx"
	xlsxTiaofu4, err := excelize.OpenFile(TiaofuPath4)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	rowsTiaofu4 := xlsxTiaofu4.GetRows(Filename4)
	//第四次计算
	Filename5 := fileNo_string+"PGa_real_4"
	TiaofuPath5 :="C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/all_seismic_tiaofu/"+Filename5+".xlsx"
	xlsxTiaofu5, err := excelize.OpenFile(TiaofuPath5)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	rowsTiaofu5 := xlsxTiaofu5.GetRows(Filename5)
	//第五次计算
	Filename6 := fileNo_string+"PGa_real_5"
	TiaofuPath6 :="C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/all_seismic_tiaofu/"+Filename6+".xlsx"
	xlsxTiaofu6, err := excelize.OpenFile(TiaofuPath6)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	rowsTiaofu6 := xlsxTiaofu6.GetRows(Filename6)

	for i := fileNO; i <= end; i++ {
		i_string := strconv.Itoa(i)
		filepath := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/DONE/"+i_string+".xlsx"
		xlsx, err := excelize.OpenFile(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rows := xlsx.GetRows("Sheet1" )
		accx,_ := strconv.ParseFloat(rows2[i][1],64)
		fmt.Println(accx)
		xishu := 100/accx
		var tiaofu []float64
		tiaofu =append(tiaofu,0,14.5,25,32,43,50,65,75,87,96,106)
		fmt.Println(xishu)
		for j:=1;j<=10;j+=1{
			j_string := strconv.Itoa(j)
			headline := "seismic_waves_"+i_string
			data := make([][]string,len(rows)+1)
			data[0] = append(data[0],headline)
			linesCount:= strconv.Itoa(len(rows)-1)
			timestep := rows[2][0]
			line2 := linesCount+ " " + timestep
			data[1] = append(data[1],line2)
			//以下是进行调幅
			for k,value := range rows{
				if k != 0{
					val,_ :=strconv.ParseFloat(value[1],64)
					tiaofu_sec,_ :=strconv.ParseFloat(rowsTiaofu[t][3],64)
					tiaofu_3rd,_ := strconv.ParseFloat(rowsTiaofu3[t][3],64)
					tiaofu_4th,_ :=strconv.ParseFloat(rowsTiaofu4[t][3],64)
					tiaofu_5th,_ :=strconv.ParseFloat(rowsTiaofu5[t][3],64)
					tiaofu_6th,_ :=strconv.ParseFloat(rowsTiaofu6[t][3],64)
					fmt.Println(tiaofu_sec,tiaofu_3rd,tiaofu_4th,tiaofu_5th,tiaofu_6th)
					value2:= fmt.Sprintf("%f",val*xishu*tiaofu[j]/5/tiaofu_sec*float64(j)/
						tiaofu_3rd*float64(j)/tiaofu_4th*float64(j)/tiaofu_5th*float64(j)/
						tiaofu_6th*float64(j))
					data[k+1] =append(data[k+1],value2)
				}
			}
			t++
			objFile :=i_string+"_"+j_string+"g.txt"
			filepath2 := "C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/Calc_File_Jacob/3-14_Jacob_214-236_371-400/"+objFile
			nfs, _:= os.OpenFile(filepath2, os.O_RDWR|os.O_CREATE, 0666)
			defer nfs.Close()
			nfs.Seek(0, io.SeekEnd)
			w := csv.NewWriter(nfs)
			//设置属性
			w.Comma = ','
			w.UseCRLF = true
			w.Flush()
			_=w.WriteAll(data)
		}
	}
}

func get_waves_1st() {
	filepath3 := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/vel_History/I_PGD_PGV.xlsx"
	xlsx2, err := excelize.OpenFile(filepath3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows2 := xlsx2.GetRows("I_PGD_PGV" )
	Filename := "237PGa_real"
	t := 1
	TiaofuPath :="C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/all_seismic_tiaofu/"+Filename+".xlsx"
	xlsxTiaofu, err := excelize.OpenFile(TiaofuPath)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	rowsTiaofu := xlsxTiaofu.GetRows(Filename)
	for i := 237; i <= 238; i++ {
		i_string := strconv.Itoa(i)
		filepath := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/DONE/"+i_string+".xlsx"
		xlsx, err := excelize.OpenFile(filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rows := xlsx.GetRows("Sheet1" )
		accx,_ := strconv.ParseFloat(rows2[i][1],64)
		fmt.Println(accx)
		xishu := 100/accx
		var tiaofu []float64
		tiaofu =append(tiaofu,0,14.5,25,32,43,50,65,75,87,96,106)
		fmt.Println(xishu)
		for j:=1;j<=10;j+=1{
			j_string := strconv.Itoa(j)
			headline := "seismic_waves_"+i_string
			data := make([][]string,len(rows)+1)
			data[0] = append(data[0],headline)
			linesCount:= strconv.Itoa(len(rows)-1)
			timestep := rows[2][0]
			line2 := linesCount+ " " + timestep
			data[1] = append(data[1],line2)
			//以下是进行调幅
			for k,value := range rows{
				if k != 0{
					val,_ :=strconv.ParseFloat(value[1],64)
					tiaofu_sec,_ :=strconv.ParseFloat(rowsTiaofu[t][3],64)
					value2:= fmt.Sprintf("%f",val*xishu*tiaofu[j]/5/tiaofu_sec*float64(j))
					data[k+1] =append(data[k+1],value2)
				}
			}
			t++
			objFile :=i_string+"_"+j_string+"g.txt"
			filepath2 := "C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/Calc_File_Jacob/2-26_Jacob_237-238/"+objFile
			nfs, _:= os.OpenFile(filepath2, os.O_RDWR|os.O_CREATE, 0666)
			defer nfs.Close()
			nfs.Seek(0, io.SeekEnd)
			w := csv.NewWriter(nfs)
			//设置属性
			w.Comma = ','
			w.UseCRLF = true
			w.Flush()
			_=w.WriteAll(data)
		}
	}
}

