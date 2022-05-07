package main

import (

	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main()  {
	var ResData [][]string
	//TODO 获取SMA（加速度时程中第三个最大值）
	for i:=1;i<=1;i++{
		i_string:= strconv.Itoa(i)
		filename := "pga/pga_seismic_"+i_string+".csv"
		time,acc:=getFile(filename)
		//获取Ars
		Ars,Tf :=getArs(time,acc)
		fmt.Printf("Ars: %f,Tf: %v\n",Ars,Tf)
		//获取Pa
		Pa:=getPa(time,Tf.tf5i,Tf.tf95i,acc)
		fmt.Printf("Pa: %f\n",Pa)
		//获取arms
		arms:= math.Sqrt(Pa)
		fmt.Printf("arms: %f\n",arms)
		//获取Ic。特征强度
		Ic := getIc(Pa,Tf.tf95,Tf.tf5)
		fmt.Printf("Ic: %f\n",Ic)
		//获取PGV
		pgvFilename := "pgv/pgv_seismic_"+i_string+".csv"
		_,pgv:=getFile(pgvFilename)
		//获取Vrms
		Vsq:= getVsq(time,pgv)
		fmt.Printf("Vsq:%f\n",Vsq)
		//获取Pv
		Pv:= getPa(time,Tf.tf5i,Tf.tf95i,pgv)
		//获取vrms
		vrms :=math.Sqrt(Pv)
		fmt.Printf("vrms:%f\n",vrms)
		//获取PGD
		pgdFilename:= "pgd/pgd_seismic_"+i_string+".csv"
		_,pgd := getFile(pgdFilename)
		Dsq:= getVsq(time,pgd)
		fmt.Printf("Dsq:%f\n",Dsq)
		//获取Pd
		Pd := getPa(time,Tf.tf5i,Tf.tf95i,pgd)
		//获取drms
		drms := math.Sqrt(Pd)
		fmt.Printf("drms: %f\n",drms)

		psaFilename:= "psa/psa_seismic_"+i_string+".csv"
		T,psa := getFile(psaFilename)
		Tg1 := 0.1*0.4//这个0.4是特征周期
		Tg2 := 0.5*0.4
		var T1,T2 int
		for j,value :=range T{
			if value == Tg1{
				T1 = j
			} else if value == Tg2{
				T2 = j
			}
		}
		ASI := getASI(T,T1,T2,psa)
		fmt.Printf("ASI:%f\n",ASI)
		Psa := getPG(psa)
		fmt.Printf("psa:%s\n",Psa)
		//获取PGv，PGD,PGA等，因为如果提前获取，会改变原时程的排序,所以放到最后
		PGV:=getPG(pgv)
		PGA :=getPG(acc)
		PGD :=getPG(pgd)
		fmt.Printf("pga:%s\n pgv:%s\n pgd:%s\n",PGA,PGV,PGD)
		//这里是获取地震动特性
		//T5_75持时
		t5_75 := Tf.tf75-Tf.tf5
		fmt.Printf("t5_75持时:%f\n",t5_75)
		//T5_95持时
		t5_95 := Tf.tf95-Tf.tf5
		fmt.Printf("t5_95持时:%f\n",t5_95)
		//一次写入多行
		Ars_string := strconv.FormatFloat(Ars,'E',-1,64)
		Pa_string := strconv.FormatFloat(Pa,'E',-1,64)
		arms_string := strconv.FormatFloat(arms,'E',-1,64)
		Ic_string := strconv.FormatFloat(Ic,'E',-1,64)
		Vsq_string := strconv.FormatFloat(Vsq,'E',-1,64)
		vrms_string := strconv.FormatFloat(vrms,'E',-1,64)
		Dsq_string := strconv.FormatFloat(Dsq,'E',-1,64)
		drms_string := strconv.FormatFloat(drms,'E',-1,64)
		ASI_string := strconv.FormatFloat(ASI,'E',-1,64)
		t5_75_string := strconv.FormatFloat(t5_75,'E',-1,64)
		t5_95_string := strconv.FormatFloat(t5_95,'E',-1,64)
		ResData = append(ResData,[]string{filename,PGA,PGV,PGD,Psa,ASI_string,Ars_string,Pa_string,arms_string,Ic_string,
			Vsq_string,vrms_string,Dsq_string,drms_string,t5_75_string,t5_95_string})
	}
	//现在进行写入文件夹的操作
	objFile := "IM_1.csv"
	NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/seismic_IM/"+objFile
	nfs, err := os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot create file,err is %+v", err)
	}
	defer nfs.Close()
	nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	objFileTitle := "IM_1"
	var title []string
	title = append(title, objFileTitle)
	err = w.Write(title)
	if err != nil {
		log.Fatalf("cannot write title：%v", err)
	}
	headline := []string{"filename","PGA", "PGV", "PGD", "Psa","ASI","Ars","Pa","Arms","Ic","Vsq","Vrms","Dsq","Drms","t5-75","t5-95"}
	err = w.Write(headline)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	//这里必须刷新，才能将数据写入文件。
	w.Flush()
	w.WriteAll(ResData)
}

func getASI(T []float64,t1 int, t2 int, psa []float64) float64{
	var asi float64
	for i:=t1;i<=t2;i++{
		result := func2_asi(T[t1],T[t1+1],psa[t1])
		asi +=math.Abs(result)
	}
	return asi
}

func func2_asi(t1 float64, t2 float64, psa float64) float64 {
	result :=math.Pow(psa,2)*(t2-t1)
	return result/2
}




func getVsq(tf []float64, pgv []float64) float64 {
	var vsq float64
	for i:=1;i<len(tf);i++{
		result := Function_dingjifen(tf[i-1],tf[i],pgv[i])
		vsq +=math.Abs(result) // 这个是总计的积分。
	}
	return vsq
}

func getPG(pg []float64) string {
	sort.Float64s(pg)
	max := pg[len(pg)-1]
	min := pg[0]
	absMax := math.Abs(min)
	if max >= absMax {
		absMax = max
	} else {
		absMax = min
	}
	pgMax_string := strconv.FormatFloat(absMax,'E',-1,64)
	return pgMax_string
}

func getIc(pa float64, tf95 float64, tf5 float64) float64{
	arms := math.Sqrt(pa)
	Ic := math.Pow(arms,1.5)*math.Pow(tf95-tf5,0.5)
	return Ic
}

func getPa(tf []float64,tf_5i,tf_95i int, acc []float64) float64{
	var jifen float64
	for i:=tf_5i;i<=tf_95i;i++{
		result := Function_dingjifen(tf[i-1],tf[i],acc[i])
		jifen +=math.Abs(result)
	}
	Pa := jifen/float64(tf_95i-tf_5i)
	return Pa
}
//定积分计算
func Function_dingjifen(t1,t2 float64,acc float64) float64 {
	result :=math.Pow(acc,3)*(t2-t1)
	return result*1/3
}
type TimeTf struct{
	jifen,jifen2 float64
	tf5,tf75,tf95 float64
	tf5i,tf75i,tf95i int
}
//计算Ars
func getArs(tf []float64, acc []float64) (float64,TimeTf){
	var Tf TimeTf
	for i:=1;i<len(tf);i++{
		result := Function_dingjifen(tf[i-1],tf[i],acc[i])
		Tf.jifen +=math.Abs(result) // 这个是总计的积分。
	}

	for i:=1; i<len(tf);i++{
		result2 := Function_dingjifen(tf[i-1],tf[i],acc[i])
		//fmt.Println(result2)
		Tf.jifen2 += math.Abs(result2)
		//这一段代码有问题。
		if Tf.jifen2<=0.05*Tf.jifen{
			Tf.tf5 = tf[i]
			Tf.tf5i = i
		}else if Tf.jifen2 >= 0.7*Tf.jifen && Tf.jifen2 <= 0.75*Tf.jifen{
			Tf.tf75 = tf[i]
			Tf.tf75i = i
		}else if Tf.jifen2 >= 0.9*Tf.jifen && Tf.jifen2 <= 0.95*Tf.jifen{
			Tf.tf95 = tf[i]
			Tf.tf95i= i
		}
	}
	Ars := math.Sqrt(Tf.jifen)
	return Ars,Tf
}
//得到PGA
func getFile(filetitle string) ([]float64,[]float64) {
	filePath := "C:/Users/JACOB/Desktop/seismic_wave/"+filetitle
	b, e := os.Open(filePath)
	if e != nil {
		fmt.Println("read file error")
		return nil,nil
	}

	reader := csv.NewReader(b)
	// 可以一次性读完
	lines, err := reader.ReadAll()
	if err != nil{
		fmt.Println("Error: ", err)
		return nil,nil
	}
	linseCount:=len(lines)
	var acc []float64
	var time []float64
	for i:=1;i<linseCount;i++{
		line := lines[i]
		timeFloat,_ :=strconv.ParseFloat(line[0],64)
		accFloat,_ :=strconv.ParseFloat(line[1],64)
		time =append(time,timeFloat)
		acc =append(acc,accFloat)
	}
	return time,acc
}



