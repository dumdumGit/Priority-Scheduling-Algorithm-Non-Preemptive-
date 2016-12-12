package main

import "fmt"

/* +-----------------------------------------------------+
/* | TUGAS SISTEM OPERASI */					      // |
/* | ALGORITMA PRIORITY SCHEDULING [NON PREEMPTIVE]      |  */
/* | 3IA07												 |  */
// +-----------------------------------------------------+


func main(){
	var burst[20] int
	var wait[20] int
	var p[20] int
	//var tat[20] int
	var pr[20] int
	var ar[20] int
	var wt[20] int

	var i,j,n,total,pos,temp,rata_wt int
	total =0

	fmt.Println()
	fmt.Printf("+-----------------------------------------------------+\n")
	fmt.Printf("| TUGAS SISTEM OPERASI     			      |\n")
	fmt.Printf("| ALGORITMA PRIORITY SCHEDULING [NON PREEMPTIVE]      |\n")
    fmt.Printf("| 3IA07                                               |\n") 
	fmt.Printf("+-----------------------------------------------------+\n\n")


	fmt.Println("Masukkan banyaknya proses : ")
	fmt.Scanln(&n)

	fmt.Printf("Masukkan Data yang dibutuhkan\n")
	for i :=0;i<n;i++{
		fmt.Printf("\n Proses P[%d] \n\n",i+1)
		fmt.Println("Burst Time : ")
		fmt.Scanln(&burst[i])
		fmt.Println("Size : ")
		fmt.Scanln(&pr[i])
		fmt.Println("Arrival Time :")
		fmt.Scanln(&ar[i])
		p[i]=i+1
	}		



for i:=1;i<n;i++{
		pos=i
		for j:=i+1;j<n;j++{

			if pr[j]>pr[pos]{
				pos=j
			}
		}
			temp = pr[i]
			pr[i] =pr[pos]
			pr[pos] = temp

			temp = burst[i]
			burst[i]= burst[pos]
		burst[pos]=temp

			temp = p[i]
			p[i] = p[pos]
			p[pos] = temp

			temp = ar[i]
			ar[i]=ar[pos]
			ar[pos]=temp

		}
	

	


	for i:=1;i<n;i++{
		
		for j=0;j<i;j++{
			wait[i] =(wait[i]+burst[j])
			wt[i] = wait[i]-ar[i]
		
		}
		total+=wt[i]
	}
	rata_wt = total/n
	total = 0


	fmt.Println()
	fmt.Println("Hasil :\n")
	fmt.Printf("+---------------------------------------------------------------------------------------------------------------+\n")
	fmt.Printf("|\tProses\t|\tBurst Time\t|\tWaiting Time\t|\tArrival Time    \t|\tSize\t|\n")
	fmt.Printf("+---------------------------------------------------------------------------------------------------------------+\n")	
	for i=0;i<n;i++{
	fmt.Printf("|\tP[%d]\t|\t%d\t\t|\t%d\t\t|\t%d\t\t\t|\t%d\t|\n",p[i],burst[i],wt[i],ar[i],pr[i])
	}
	fmt.Printf("+---------------------------------------------------------------------------------------------------------------+\n")
	//rata_tat = total/n
	fmt.Printf("\n Average Waiting Time is:%d",rata_wt)
	/*fmt.Printf("\n Rata-rata Turn Around Time adalah:%d",rata_tat) */
	fmt.Println("\n\n")

	for i:=0;i<n;i++{
		fmt.Printf("  P[%d]    ",p[i])
	}

	fmt.Println()

	for i:=0;i<n;i++{
		fmt.Printf("|")
		fmt.Printf("========|")
	}
	fmt.Println()
	for i:=0;i<n;i++{

		fmt.Printf("%d        ",wait[i])
	}
	fmt.Println()
}
