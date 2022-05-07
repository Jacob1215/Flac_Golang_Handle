package main

import (
	"encoding/csv"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"strconv"
)

func main () {
	//call_1st_accx_only()
	get()
}

func call_1st_accx_only() {
	rows := make([][]string,13)
	rows[0] = append(rows[0],";seismic_waves_1-1g")
	rows[1] = append(rows[1],"restore 145_2.sav")
	rows[2] = append(rows[2],"initial xdisp 0 ydisp 0")
	rows[3] = append(rows[3],"initial xvel 0 yvel 0")
	rows[4] = append(rows[4],"hist read 1_1g.txt")
	rows[5] = append(rows[5],"call 'D:/Jacob/CALL_FILE/apply_range.txt'")
	rows[6] = append(rows[6],"set step=100000000")
	rows[7] = append(rows[7],"set dyn=on")
	rows[8] = append(rows[8],"set dy_damping rayleigh=0.0384 15.38")
	rows[9] = append(rows[9],"solve dytime 26.845")
	rows[10] = append(rows[10],"set hisfile=seismic_1_0.1_accx.txt")
	rows[11] = append(rows[11],"hist write 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 vs 1 skip 5")
	filepath2 := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/vel_History/I_PGD_PGV.xlsx"
	xlsx2, _ := excelize.OpenFile(filepath2)
	rows2 := xlsx2.GetRows("I_PGD_PGV" )
	filepath3 := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt_pga/new/RayleiZuni.xlsx"
	xlsx3, _ := excelize.OpenFile(filepath3)
	rows3 := xlsx3.GetRows("RayleiZuni" )
	title:= "accx_only_seismic_waves_237-238_2-26"
	for i :=214; i<=238; i++{//gut,就这样
		i_string :=strconv.Itoa(i)
		for j:=1;j<=10;j++{
			j_string := strconv.Itoa(j)
			lines := rows
			lines[0][0] = ";seismic_waves_"+i_string+"_"+j_string+"g"
			lines[4][0] = "hist read "+i_string+"_"+j_string+"g.txt"
			lines[8][0] = "set dy_damping rayleigh="+rows3[i][1]+" "+rows3[i][2]
			lines[9][0] = "solve dytime "+rows2[i][4]
			lines[10][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_accx.txt"
			NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/Calc_File_Jacob/"+title+"_accx_only.txt"
			nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
			defer nfs.Close()
			nfs.Seek(0, io.SeekEnd)
			w := csv.NewWriter(nfs)
			//设置属性
			w.Comma = ','
			w.UseCRLF = true
			w.Flush()
			_=w.WriteAll(lines)
		}
	}
}


func get() {
	rows := make([][]string,28)
	rows[0] = append(rows[0],";seismic_waves_1-1g")
	rows[1] = append(rows[1],"restore 145_2.sav")
	rows[2] = append(rows[2],"initial xdisp 0 ydisp 0")
	rows[3] = append(rows[3],"initial xvel 0 yvel 0")
	rows[4] = append(rows[4],"hist read 1_1g.txt")
	rows[5] = append(rows[5],"call 'D:/Jacob/CALL_FILE/apply_range.txt'")
	rows[6] = append(rows[6],"set step=100000000")
	rows[7] = append(rows[7],"set dyn=on")
	rows[8] = append(rows[8],"set dy_damping rayleigh=0.0384 15.38")
	rows[9] = append(rows[9],"solve dytime 26.845")
	rows[10] = append(rows[10],"set hisfile=seismic_1_0.1_accx.txt")
	rows[11] = append(rows[11],"hist write 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 vs 1 skip 5")
	rows[12] = append(rows[12],"set hisfile=seismic_1_0.1_moment.txt")
	rows[13] = append(rows[13],"hist write 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 vs 1 skip 5")
	rows[14] = append(rows[14],"set hisfile=seismic_1_0.1_soil_xdisp.txt")
	rows[15] = append(rows[15],"hist write 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 vs 1 skip 5")
	rows[16] = append(rows[16],"set hisfile=seismic_1_0.1_xvel.txt")
	rows[17] = append(rows[17],"hist write 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 vs 1 skip 5")
	rows[18] = append(rows[18],"set hisfile=seismic_1_0.1_tunnel_xdisp.txt")
	rows[19] = append(rows[19],"hist write 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 vs 1 skip 5")
	rows[20] = append(rows[20],"set hisfile=seismic_1_0.1_tunnel_ydisp.txt")
	rows[21] = append(rows[21],"hist write 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128 129 130 131 132 133 134 135 136 137 138 139 140 141 142 vs 1 skip 5")
	rows[22] = append(rows[22],"set hisfile=seismic_1_0.1_axial.txt")
	rows[23] = append(rows[23],"hist write 144 145 146 147 148 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168 169 170 171 172 173 174 175 vs 1 skip 5")
	rows[24] = append(rows[24],"set hisfile=seismic_1_0.1_shear.txt")
	rows[25] = append(rows[25],"hist write 176 177 178 179 180 181 182 183 184 185 186 187 188 189 190 191 192 193 194 195 196 197 198 199 200 201 202 203 204 205 206 207 vs 1 skip 5")
	rows[26] = append(rows[26],"save seismic_waves_1-1g.sav")
	filepath2 := "C:/Users/JACOB/Desktop/seismic_wave/getPGA/vel_History/I_PGD_PGV.xlsx"
	xlsx2, _ := excelize.OpenFile(filepath2)
	rows2 := xlsx2.GetRows("I_PGD_PGV" )
	filepath3 := "C:/Users/JACOB/Desktop/seismic_wave/Flac_code/waves_txt_pga/new/RayleiZuni.xlsx"
	xlsx3, _ := excelize.OpenFile(filepath3)
	rows3 := xlsx3.GetRows("RayleiZuni" )
	title:= "seismic_waves_3-13_1-12"
	for i :=1; i<=12; i++{//gut,就这样
		i_string :=strconv.Itoa(i)
		for j:=1;j<=10;j++{
			j_string := strconv.Itoa(j)
			lines := rows
			lines[0][0] = ";seismic_waves_"+i_string+"_"+j_string+"g"
			lines[4][0] = "hist read "+i_string+"_"+j_string+"g.txt"
			lines[8][0] = "set dy_damping rayleigh="+rows3[i][1]+" "+rows3[i][2]
			lines[9][0] = "solve dytime "+rows2[i][4]
			lines[10][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_accx.txt"
			lines[12][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_moment.txt"
			lines[14][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_soil_xdisp.txt"
			lines[16][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_xvel.txt"
			lines[18][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_tunnel_xdisp.txt"
			lines[20][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_tunnel_ydisp.txt"
			lines[22][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_axial.txt"
			lines[24][0] = "set hisfile=seismic_"+i_string+"_0."+j_string+"_shear.txt"
			lines[26][0] = "save seismic_waves_"+i_string+"_"+j_string+"g.sav"
			NewFileName := "C:/Users/JACOB/Desktop/seismic_wave/2_15_Re_Calc/Calc_File_Jacob/"+title+".txt"
			nfs, _:= os.OpenFile(NewFileName, os.O_RDWR|os.O_CREATE, 0666)
			defer nfs.Close()
			nfs.Seek(0, io.SeekEnd)
			w := csv.NewWriter(nfs)
			//设置属性
			w.Comma = ','
			w.UseCRLF = true
			w.Flush()
			_=w.WriteAll(lines)
		}
	}

}