package main

import (
	"database/sql"
	"os"
//	"bufio"
	"bytes"
	"flag"
	"log"
	"net/http"
	"strings"
	"strconv"
	"time"
	"github.com/gorilla/websocket"
	_ "github.com/go-sql-driver/mysql" )

func dev(){
	v := strconv.Itoa(20); log.Println(v)
//	v := max(2,3)
}

/*
type Cq struct{ 
	mas, ie, s 			bool
	c 					int8
	ni 					string}
type Rq struct{
	Ac, SN 				bool
	F, T 				int8
	ni, desc 			string}
type Cl struct{ 
	hue, sat, val, c, r, gamma, beta 	int8
	hex 				string}

type Ml struct{
	ni, desc 	string}
*/

func abs(n int8) int8 { 	if n < 0 { return -n} else { return n}}
//func max(n1 int8, n2 int8){ if n1<n2 { return n2} else { return n1}}
func min(n1 uint8, n2 uint8) uint8 { if n1>n2 { return n2} else { return n1}}
func odd(n int8) bool { 	if n%2==0{ return false} else { return true}}


type Va struct{ // ööl
	inV, inK, exV, exK 	int}




type People struct{}
type Family struct{}
type Goal struct{}
type Model struct{}
type Subj struct{
	gns 				[4]Gn
	pms 				[12]Pm
	mds 				[]Md
	mp 					Mp
	xtn,ntn 			Pt
	fn 					Fn
	cr 					Cr
	tp 					Tp
	}
type Obj struct{}
type Dharma struct{}
type Duty struct{}


type Sim struct{
	copy 	bool
	src 	uint
	gNs [4]Gn; 	pMs [12]Pm;	mDs [192]Md;cQs [40]Cq;	iIs [8]Ii;fNs [8]Fn;cRs [12]Cr;	tPs [16]Tp;	iRs [16]Ir;	rQs [40]Rq
	gns []Gn;	pms []Pm;	mds []Md; 	cqs []Cq;	iis []Ii; fns []Fn;	crs []Cr;	tps []Tp;	irs []Ir;	rqs []Rq
	peoples []People;	familys []Family;	goals 	[]Goal;		models 	[]Model
	subjs 	[]Subj;		objs 	[]Obj;		dharmas []Dharma;	duties 	[]Duty}
func (s Sim) copyMetaphysics(	gns [4]Gn, pms [12]Pm, mds [192]Md, cqs [40]Cq, iis [8]Ii, fns [8]Fn, crs [12]Cr, tps [16]Tp, irs [16]Ir, rqs [40]Rq){
								s.gNs=gns; s.pMs=pMs;  s.mDs=mds;   s.cQs=cqs;  s.iIs=iis; s.fNs=fns; s.cRs=crs;  s.tPs=tps;  s.iRs=irs;  s.rQs=rqs}

type Stat struct{
	avg, med 			bool
}


type User struct{
	sim 	*Sim
	sims 	[]Sim
	simsS 	[]*Sim
	stat 	*Stat
	gNs [4]Gn; 	pMs [12]Pm;	mDs [192]Md;cQs [40]Cq;	iIs [8]Ii;fNs [8]Fn;cRs [12]Cr;	tPs [16]Tp;	iRs [16]Ir;	rQs [40]Rq
	gns []Gn;	pms []Pm;	mds []Md; 	cqs []Cq;	iis []Ii; fns []Fn;	crs []Cr;	tps []Tp;	irs []Ir;	rqs []Rq
	peoples []People;	familys []Family;	goals 	[]Goal;		models 	[]Model
	subjs 	[]Subj;		objs 	[]Obj;		dharmas []Dharma;	duties 	[]Duty}
func (u User) copyMetaphysics(	gns [4]Gn, pms [12]Pm, mds [192]Md, cqs [40]Cq, iis [8]Ii, fns [8]Fn, crs [12]Cr, tps [16]Tp, irs [16]Ir, rqs [40]Rq){
								u.gNs=gns; u.pMs=pMs;  u.mDs=mds;   u.cQs=cqs;  u.iIs=iis; u.fNs=fns; u.cRs=crs;  u.tPs=tps;  u.iRs=irs;  u.rQs=rqs}
func (u User) mkSim(copy bool,src uint) { var s Sim;
	s.copy =copy; if copy { s.src =src}
	s.copyMetaphysics( gNs, pMs, mDs, cQs, iIs, fNs, cRs, tPs, iRs, rQs)
	u.sims =append(u.sims, s)}
func initUser(){
	if len(users)==0 { mkUser()}}
func mkUser(){ var u User;
 	u.copyMetaphysics( gNs, pMs, mDs, cQs, iIs, fNs, cRs, tPs, iRs, rQs)
 	//u.mkSim(false,0,pcPeoples,pcFamilys,pcGoals,pcModels,pcSubjs,pcObjs,pcDharmas,pcDuties)
 	users =append(users, u)}
//subjs, objs, dharmas, duties, ppls, famis, goals, models
var users []User



/*sound like clicking the keyboard is associated to a vector quality*/




func initEssences(){
	}
var defPeoples 	[]People
var defFamilys 	[]Family
var defGoals 	[]Goal
var defModels 	[]Model
var defSubjs 	[]Subj
var defObjs 	[]Obj
var defDharmas 	[]Dharma
var defDuties 	[]Duty
//var users = []

// if no subj is selected then can select multiple types or moods

//func mkSim(u User) Sim {

//}

/*type Test struct{
	att 	int
}*/












type Rq struct{
	Ac,sn 				bool
	F,T 				int8}
// mod array for models, empty in archetype
//	lpft,lpsn/*pps,ppo*/Va
//	pts 				[]Pt}
func specRq(	Ac bool, sn bool, 	F int8, T int8) Rq { var rq Rq
				rq.Ac=Ac;rq.sn=sn; 	rq.F=F; rq.T=T; return rq}
func precalcRomanticQualities(){
	rQs[0]=specRq( false,false, 0,0); 	rQs[20]=specRq(true ,false, 0,0);
	rQs[1]=specRq( false,false, 1,4); 	rQs[21]=specRq(true ,false, 1,4);
	rQs[2]=specRq( false,false, 2,3); 	rQs[22]=specRq(true ,false, 2,3);
	rQs[3]=specRq( false,false, 3,2); 	rQs[23]=specRq(true ,false, 3,2);
	rQs[4]=specRq( false,false, 4,1); 	rQs[24]=specRq(true ,false, 4,1);

	rQs[5]=specRq( false,false, 0,0); 	rQs[25]=specRq(true ,false, 0,0);
	rQs[6]=specRq( false,false, 1,4); 	rQs[26]=specRq(true ,false, 1,4);
	rQs[7]=specRq( false,false, 2,3); 	rQs[27]=specRq(true ,false, 2,3);
	rQs[8]=specRq( false,false, 3,2); 	rQs[28]=specRq(true ,false, 3,2);
	rQs[9]=specRq( false,false, 4,1); 	rQs[29]=specRq(true ,false, 4,1);

	rQs[10]=specRq(false,true , 0,0); 	rQs[30]=specRq(true ,true , 0,0);
	rQs[11]=specRq(false,true , 1,4); 	rQs[31]=specRq(true ,true , 1,4);
	rQs[12]=specRq(false,true , 2,3); 	rQs[32]=specRq(true ,true , 2,3);
	rQs[13]=specRq(false,true , 3,2); 	rQs[33]=specRq(true ,true , 3,2);
	rQs[14]=specRq(false,true , 4,1); 	rQs[34]=specRq(true ,true , 4,1);

	rQs[15]=specRq(false,true , 0,0); 	rQs[35]=specRq(true ,true , 0,0);
	rQs[16]=specRq(false,true , 1,4); 	rQs[36]=specRq(true ,true , 1,4);
	rQs[17]=specRq(false,true , 2,3); 	rQs[37]=specRq(true ,true , 2,3);
	rQs[18]=specRq(false,true , 3,2); 	rQs[38]=specRq(true ,true , 3,2);
	rQs[19]=specRq(false,true , 4,1); 	rQs[39]=specRq(true ,true , 4,1)}
var rQs [40]Rq


type Cq struct{
	mas,dir,ft,ie,sn,ws	bool
	mag			int8}
func specCq(	mas bool, 	ft bool, 	mag int8  	) Cq { var cq Cq;
				cq.mas=mas; cq.ft=ft; 	cq.mag=mag;
	if !ft { if mag <3{ cq.ie =false} else{ cq.ie =true }; 	if odd(mag) { cq.ws =false} else{ cq.ws =true }
	} else { if mag <3{ cq.ie =true } else{ cq.ie =false}; 	if odd(mag) { cq.ws =true } else{ cq.ws =false}}; return cq}
func precalcClassicalQualities(){ 
	cQs[0]=specCq( false,false, 0); 	cQs[20]=specCq(true ,false, 0);
	cQs[1]=specCq( false,false, 1); 	cQs[21]=specCq(true ,false, 1);
	cQs[2]=specCq( false,false, 2); 	cQs[22]=specCq(true ,false, 2);
	cQs[3]=specCq( false,false, 3); 	cQs[23]=specCq(true ,false, 3);
	cQs[4]=specCq( false,false, 4); 	cQs[24]=specCq(true ,false, 4);

	cQs[5]=specCq( false,false, 0); 	cQs[25]=specCq(true ,false, 0);
	cQs[6]=specCq( false,false, 1); 	cQs[26]=specCq(true ,false, 1);
	cQs[7]=specCq( false,false, 2); 	cQs[27]=specCq(true ,false, 2);
	cQs[8]=specCq( false,false, 3); 	cQs[28]=specCq(true ,false, 3);
	cQs[9]=specCq( false,false, 4); 	cQs[29]=specCq(true ,false, 4);

	cQs[10]=specCq(false,true , 0); 	cQs[30]=specCq(true ,true , 0);
	cQs[11]=specCq(false,true , 1); 	cQs[31]=specCq(true ,true , 1);
	cQs[12]=specCq(false,true , 2); 	cQs[32]=specCq(true ,true , 2);
	cQs[13]=specCq(false,true , 3); 	cQs[33]=specCq(true ,true , 3);
	cQs[14]=specCq(false,true , 4); 	cQs[34]=specCq(true ,true , 4);

	cQs[15]=specCq(false,true , 0); 	cQs[35]=specCq(true ,true , 0);
	cQs[16]=specCq(false,true , 1); 	cQs[36]=specCq(true ,true , 1);
	cQs[17]=specCq(false,true , 2); 	cQs[37]=specCq(true ,true , 2);
	cQs[18]=specCq(false,true , 3); 	cQs[38]=specCq(true ,true , 3);
	cQs[19]=specCq(false,true , 4); 	cQs[39]=specCq(true ,true , 4)}
var cQs [40]Cq



type Ir struct{
	cr, ie, ft, sn 		bool}
func specIr(cr bool, 	ie bool, 	ft bool, 	sn bool) Ir { var ir Ir
			ir.cr=cr; 	ir.ie=ie; 	ir.ft=ft; 	ir.sn=sn; return ir}
func precalcIntertypeRelations(){
	iRs[0] =specIr(false,false,false,false)
	iRs[1] =specIr(true, false,false,false)
	iRs[2] =specIr(false,true, false,false)
	iRs[3] =specIr(true, true, false,false)
	iRs[4] =specIr(false,false,true, false)
	iRs[5] =specIr(true, false,true, false)
	iRs[6] =specIr(false,true, true, false)
	iRs[7] =specIr(true, true, true, false)
	iRs[8] =specIr(false,false,false,true )
	iRs[9] =specIr(true, false,false,true )
	iRs[10]=specIr(false,true, false,true )
	iRs[11]=specIr(true, true, false,true )
	iRs[12]=specIr(false,false,true, true )
	iRs[13]=specIr(true, false,true, true )
	iRs[14]=specIr(false,true, true, true )
	iRs[15]=specIr(true, true, true, true)}
var iRs [16]Ir;



type Tp struct{
	cr, ie, ft, sn 		bool
	iis 				[8]Ii
	fns 				[8]Fn
	crs 				[12]Cr
	ii3a,ii3r			Ii3d
	fn3a,fn3r 			Fn3d}
type Ii3d [2][2][2]*Ii
func (m Ii3d) trans3d(t uint8) Ii3d { switch t { case 1:m=m.lon(); case 2:m=m.lat(); case 3:m=m.ver(); case 4:m=m.rot(); case 5:m=m.tor()}; return m}
func (m Ii3d) lon() Ii3d {var M Ii3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[1-z][y][x] 	}}}; return M}
func (m Ii3d) lat() Ii3d {var M Ii3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][1-y][x] 	}}}; return M}
func (m Ii3d) ver() Ii3d {var M Ii3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][y][1-x] 	}}}; return M}
func (m Ii3d) rot() Ii3d {var M Ii3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][x][y] 	}}}; return M}
func (m Ii3d) tor() Ii3d {var M Ii3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][1-y][1-x] }}}; return M}
func mkIi3d(iis [8]Ii) Ii3d { var m Ii3d;
    //sn,ie,ft
	m[0][0][0] =&iis[0]; 			//Fi
	m[1][0][0] =&iis[1]; 			//Ni
	m[0][1][0] =&iis[2]; 			//Fe
	m[1][1][0] =&iis[3]; 			//Ne
	m[0][0][1] =&iis[4]; 			//Si
	m[1][0][1] =&iis[5]; 			//Ti
	m[0][1][1] =&iis[6]; 			//Se
	m[1][1][1] =&iis[7]; return m} 	//Te
type Fn3d [2][2][2]*Fn
func (m Fn3d) trans3d(t uint8) Fn3d { switch t { case 1:m=m.lon(); case 2:m=m.lat(); case 3:m=m.ver(); case 4:m=m.rot(); case 5:m=m.tor()}; return m}
func (m Fn3d) lon() Fn3d {var M Fn3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[1-z][y][x] 	}}}; return M}
func (m Fn3d) lat() Fn3d {var M Fn3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][1-y][x] 	}}}; return M}
func (m Fn3d) ver() Fn3d {var M Fn3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][y][1-x] 	}}}; return M}
func (m Fn3d) rot() Fn3d {var M Fn3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][x][y] 	}}}; return M}
func (m Fn3d) tor() Fn3d {var M Fn3d;for z:=int8(0);z<2;z++{for y:=int8(0);y<2;y++{for x:=int8(0);x<2;x++{M[z][y][x]=m[z][1-y][1-x] }}}; return M}
func mkFn3d(fns [8]Fn) Fn3d { var m Fn3d;
    //verlatlon
	m[0][0][0] =&fns[0]; 
	m[1][0][0] =&fns[1]; 
	m[0][1][0] =&fns[2]; 
	m[1][1][0] =&fns[3]; 
	m[0][0][1] =&fns[4]; 
	m[1][0][1] =&fns[5]; 
	m[0][1][1] =&fns[6]; 
	m[1][1][1] =&fns[7]; return m}

func (tp Tp) trans3d(t1 uint8, t2 uint8, t3 uint8) (Fn3d,Ii3d) { fn3a:=tp.fn3a; ii3r:=tp.ii3r
				fn3a=fn3a.trans3d(t1); 	ii3r=ii3r.trans3d(t1);
	if t2 >0{ 	fn3a=fn3a.trans3d(t2); 	ii3r=ii3r.trans3d(t2)}
	if t3 >0{ 	fn3a=fn3a.trans3d(t3); 	ii3r=ii3r.trans3d(t3)}
	return fn3a, ii3r}
func defCrs(crs [12]Cr, fns[8]Fn) [12]Cr { a:=crs
	a[0].fn =&fns[0]
	a[1].fn =&fns[2]
	a[2].fn =&fns[4]
	a[3].fn =&fns[6]
	a[4].fn =&fns[4]
	a[5].fn =&fns[0]
	a[6].fn =&fns[1]
	a[7].fn =&fns[5]
	a[8].fn =&fns[3]
	a[9].fn =&fns[5]
	a[10].fn=&fns[6]
	a[11].fn=&fns[0];return a}
func defFns(f Fn3d, i Ii3d) Fn3d {
	f[0][0][0].ii =i[0][0][0]
	f[0][0][1].ii =i[0][0][1]
	f[0][1][0].ii =i[0][1][0]
	f[0][1][1].ii =i[0][1][1]
	f[1][0][0].ii =i[1][0][0]
	f[1][0][1].ii =i[1][0][1]
	f[1][1][0].ii =i[1][1][0]
	f[1][1][1].ii =i[1][1][1];return f}
func specTp(cr bool, 	ie bool, 	ft bool, 	sn bool) Tp { var tp Tp
			tp.cr=cr; 	tp.ie=ie; 	tp.ft=ft; 	tp.sn=sn;
	tp.iis =iIs; 	tp.ii3a =mkIi3d(tp.iis); 	tp.ii3r =mkIi3d(tp.iis);
	tp.fns =fNs; 	tp.fn3a =mkFn3d(tp.fns); 	tp.fn3r =mkFn3d(tp.fns);
	if !sn{ if !ft{ if !ie{ if !cr { tp.fn3a, tp.ii3r =tp.trans3d(2,3,0)	//iSFC lon1 lat2 ver3 rot4 tor5 
							} else { tp.fn3a, tp.ii3r =tp.trans3d(1,2,3)
					}}else {if !cr { tp.fn3a, tp.ii3r =tp.trans3d(3,0,0)
							} else { tp.fn3a, tp.ii3r =tp.trans3d(1,2,0)
			}}}else{if !ie{ if !cr { tp.fn3a, tp.ii3r =tp.trans3d(2,5,0)
							} else { tp.fn3a, tp.ii3r =tp.trans3d(1,2,5)
					}}else {if !cr { tp.fn3a, tp.ii3r =tp.trans3d(5,0,0)
							} else { tp.fn3a, tp.ii3r =tp.trans3d(1,5,0)
	}}}}else{if !ft{if !ie{ if !cr { tp.fn3a, tp.ii3r =tp.trans3d(2,4,0)
							} else { tp.fn3a, tp.ii3r =tp.trans3d(1,2,4)
					}}else {if !cr { tp.fn3a, tp.ii3r =tp.trans3d(1,4,0)
							} else { tp.fn3a, tp.ii3r =tp.trans3d(4,0,0)
			}}}else{if !ie{ if !cr { tp.fn3a, tp.ii3r =tp.trans3d(1,2,0)
							} else { tp.fn3a, tp.ii3r =tp.trans3d(2,0,0)
					}}else {if !cr { tp.fn3a, tp.ii3r =tp.trans3d(1,0,0)
							} else { }}}} 									//eNTR
	tp.fn3a=defFns( tp.fn3a, tp.ii3a)
	tp.crs =defCrs( cRs, tp.fns);
	return tp}
func precalcSociotypes(){
	tPs[0] =specTp( false,false,false,false)
	tPs[1] =specTp( true, false,false,false)
	tPs[2] =specTp( false,true, false,false)
	tPs[3] =specTp( true, true, false,false)
	tPs[4] =specTp( false,false,true, false)
	tPs[5] =specTp( true, false,true, false)
	tPs[6] =specTp( false,true, true, false)
	tPs[7] =specTp( true, true, true, false)
	tPs[8] =specTp( false,false,false,true )
	tPs[9] =specTp( true, false,false,true )
	tPs[10]=specTp( false,true, false,true )
	tPs[11]=specTp( true, true, false,true )
	tPs[12]=specTp( false,false,true, true )
	tPs[13]=specTp( true, false,true, true )
	tPs[14]=specTp( false,true, true, true )
	tPs[15]=specTp( true, true, true, true)}
var tPs [16]Tp




type Cr struct{
	ws, vm 				bool
	res,spd				int8
	lp 					Lp
	pp 					Pp
	fn 					*Fn}
func specCr(ws bool, vm bool,  	res int8) Cr { var cr Cr
			cr.ws=ws;cr.vm=vm; 	cr.res=res; return cr }
func precalcCircuits(){
//	cRs[0] =specCr( false,false,  0)
	cRs[0] =specCr( true ,true , -2); 	cRs[0].lp =specLp(-1,-1); 	cRs[0].pp =specPp(-1,-1) 	//love
	cRs[1] =specCr( false,false,  2); 	cRs[1].lp =specLp(-1, 1); 	cRs[1].pp =specPp(-1, 1) 	//dare
	cRs[2] =specCr( false,true , -2); 	cRs[2].lp =specLp( 1, 1); 	cRs[2].pp =specPp( 1, 1) 	//chal
	cRs[3] =specCr( true ,false,  2); 	cRs[3].lp =specLp( 1,-1); 	cRs[3].pp =specPp( 1,-1) 	//harv
	cRs[4] =specCr( true ,true ,  1); 	cRs[4].lp =specLp(-1,-1); 	cRs[4].pp =specPp(-1,-1) 	//ego
	cRs[5] =specCr( false,true , -1); 	cRs[5].lp =specLp(-1, 1); 	cRs[5].pp =specPp( 1,-1) 	//sego
	cRs[6] =specCr( false,false, -1); 	cRs[6].lp =specLp( 1, 1); 	cRs[6].pp =specPp( 1,-1) 	//sid
	cRs[7] =specCr( true ,false,  1); 	cRs[7].lp =specLp( 1,-1); 	cRs[7].pp =specPp(-1, 1) 	//id
	cRs[8] =specCr( true ,false,  3); 	cRs[8].lp =specLp(-1,-1); 	cRs[8].pp =specPp(-1,-1) 	//neur
	cRs[9] =specCr( false,true , -3); 	cRs[9].lp =specLp( 1,-1); 	cRs[9].pp =specPp(-1, 1) 	//meta
	cRs[10]=specCr( true ,true , -3); 	cRs[10].lp=specLp(-1, 1); 	cRs[10].pp=specPp( 1,-1) 	//morp
	cRs[11]=specCr( true ,false,  3); 	cRs[11].lp=specLp( 1, 1); 	cRs[11].pp=specPp( 1, 1)}	//mego
//	cRs[13]=specCr( true ,true ,  4)
var cRs [12]Cr



type Fn struct{
	lon,ver,lat 	bool
	ii 				*Ii}
func specFn( 	lon bool, 	lat bool, 	ver bool) Fn { var fn Fn;
				fn.lon=lon;	fn.lat=lat;	fn.ver=ver; return fn}
func precalcFunctions(){
	fNs[2]=specFn(false,false,false);//mobi
	fNs[5]=specFn(false,false,true );//igno
	fNs[3]=specFn(false,true, false);//prov
	fNs[4]=specFn(false,true, true );//ego
	fNs[1]=specFn(true, false,false);//suggest
	fNs[6]=specFn(true, false,true );//demo
	fNs[0]=specFn(true, true, false);//role
	fNs[7]=specFn(true, true, true )}//crea
var fNs [8]Fn



type Ii struct{
	ft,ie,sn, dg  	bool}
func specIi( 	ft bool, ie bool, sn bool, 	dg bool) Ii { var ii Ii;
				ii.ft=ft;ii.ie=ie;ii.sn=sn;	ii.dg=dg; return ii}
func precalcInformationIsotopes(){
	iIs[7]=specIi(false,false,false, false)	//Fi
	iIs[1]=specIi(false,false,true,  true )	//Ni
	iIs[6]=specIi(false,true, false, true )	//Fe
	iIs[0]=specIi(false,true, true,  false)	//Ne
	iIs[5]=specIi(true, false,false, true )	//Si
	iIs[3]=specIi(true, false,true,  false)	//Ti
	iIs[4]=specIi(true, true, false, false) //Se
	iIs[2]=specIi(true, true, true,  true )}//Te
var iIs [8]Ii



type Mp struct{
	lx,ly,dx,dy 		int8
	ul,ur,dr,dl 		Qp
	mm 					[18][18]Mv}
type Mv struct{
	ide 				int8
	col 				Color
	qv 					Qv}
type Pt struct{
	cq 					Cq
	rq 					Rq
	mv 					Mv}


/*func pointCqByIi( cqs1 []CqS, ft bool,ie bool,sn bool) []*Cq { var cqs2 []*Cq;
	for i := int8(0); i <40; i++{ if cqs1.ft==ft && cqs1.ie==ie && cqs1.sn==sn{ cqs2 =append(cqs2, &cqs1[i])}}; return cqs2}*/


type Md struct{
	sx, gn 				*Gn
	pm 					*Pm
	gwf, bwf 			Wf
	gar, bar 			Awf
	qp 					Qp}
func (md Md) modPm(pm *Pm) (Wf,Wf) {
	md.gwf =md.sx.gwf.modPm2(pm.gwf); 	md.bwf =md.sx.bwf.modPm2(pm.bwf); return md.gwf, md.bwf}
func (f1 Wf) modPm2(pm Wf) Wf { var f2 Wf
	f2[0] =f1[pm[0]]; 	f2[1] =f1[pm[1]]; 	f2[2] =f1[pm[2]]; 	return f2}
func (md Md) modMd(t1 int8, t2 int8) (Wf,Wf) {
	if t1 > 0 {	for i := int8(0); i < 3; i++ { var f Wf; f =md.gwf.modWf(0,1,2);
		if md.gwf[i] ==t1 { f[i] =t2}; 	if md.gwf[i] ==t2 { f[i] =t1}; return f, md.bwf}
	} else { 	for i := int8(0); i < 3; i++ { var f Wf; f =md.bwf.modWf(0,1,2);
		if md.bwf[i] ==t1 { f[i] =t2}; 	if md.bwf[i] ==t1 { f[i] =t1}; return md.gwf, f}}
	return md.gwf, md.bwf}
type Awf func(int8,int8) (int8,int8)
type Qp [9][9]Qv
type Qv struct{
	c, r 				int8
	lp 					Lp}
func precalcMoods(){
	for i := uint8(0); i < 48; i++ {
		for j := uint8(0); j < 4;  j++ {
			mDs[i+48*j].sx =&gNs[j]}}
	for i := uint8(0); i < 12; i++ {
		for j := uint8(0); j < 16; j++ {
			mDs[i+12*j].pm =&pMs[i];
			switch j {
			case 0, 4, 8, 12: mDs[i+12*j].gn =&gNs[0];
			case 1, 5, 9, 13: mDs[i+12*j].gn =&gNs[1];
			case 2, 6, 10,14: mDs[i+12*j].gn =&gNs[2];
			case 3, 7, 11,15: mDs[i+12*j].gn =&gNs[3]}}}
	for i := uint8(0); i < 192;i++ {
		mDs[i].gwf, mDs[i].bwf =mDs[i].modPm( mDs[i].pm);
		if mDs[i].sx == mDs[i].gn { 													// if sex and gender aren't same 		
			if mDs[i].sx.mas != mDs[i].gn.mas && mDs[i].sx.Ac == mDs[i].gn.Ac { 		// .if sex and gender aren't foam 		
				if !mDs[i].sx.mas { mDs[i].gwf, mDs[i].bwf =mDs[i].modMd(3,-4) 		// ..if sex ain't masculine...good 3 bad -4 f_iQUEER-
				} else { 			mDs[i].gwf, mDs[i].bwf =mDs[i].modMd(3,-5)} 		// ..otherwise good 3 bad -5 			m_eQUEER-
			} else if mDs[i].sx.mas == mDs[i].gn.mas && mDs[i].sx.Ac != mDs[i].gn.Ac { 	// .else if sex and gender aren't time 	
				if !mDs[i].sx.Ac { 	mDs[i].gwf, mDs[i].bwf =mDs[i].modMd(3,-6) 		// ..if sex ain't time...good 3 bad -6 		A_AGE-
				} else { 			mDs[i].gwf, mDs[i].bwf =mDs[i].modMd(-3,6)} 		// ..otherwise good -3 bad 6 			c_AGE+
			} else { 																	// if sex and gender are the same 		
				mDs[i].gwf, mDs[i].bwf =mDs[i].modMd(3,-6); 							// .good 3 bad -6 						f_iQUEER+
				mDs[i].gwf, mDs[i].bwf =mDs[i].modMd(-3,6)}} 							// .and good -3 bad 6 					m_eQUEER+
		mDs[i].gar=mDs[i].gwf.arithWf()
		mDs[i].bar=mDs[i].bwf.arithWf()
		mDs[i].qp =mDs[i].quantQp( mDs[i].gar,mDs[i].bar)}}
func (wf Wf) arithWf() Awf { var a Awf
	if 			wf[0] == 1 {if wf[1] == 2 { if wf[2] == 3 {	a =func(c int8, r int8)(int8,int8){ return 	-r,	c-r} 	// foam phi alpha psi beta rho gamma 	: 
									} else 	if wf[2] ==-4 {	a =func(c int8, r int8)(int8,int8){ return 	-r,	c-r} 	// 							rho -4
									} else 	if wf[2] ==-5 {	a =func(c int8, r int8)(int8,int8){ return 	-r,	c-r} 	// 							rho -5
									} else 	{ 				a =func(c int8, r int8)(int8,int8){ return 	-r,	c-r}} 	// 							rho -6
					} else 	if wf[2] == 2 { if wf[1] == 3 { a =func(c int8, r int8)(int8,int8){ return c-r, -c} 	// 					psi beta rho gamma
									} else	if wf[1] ==-4 { a =func(c int8, r int8)(int8,int8){ return c-r, -c} 	// 					psi -4
									} else 	if wf[1] ==-5 { a =func(c int8, r int8)(int8,int8){	return c-r, -c} 	// 					psi -5
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return c-r, -c}}} 	// 					psi -6
	} else 	if wf[0] == 2 {	if wf[1] == 1 { if wf[2] == 3 { a =func(c int8, r int8)(int8,int8){	return 	-r,	4-c} 	// 		phi beta psi alpha rho gamma
									} else 	if wf[2] ==-4 { a =func(c int8, r int8)(int8,int8){	return 	-r,	4-c} 	// 							rho -4
									} else 	if wf[2] ==-5 { a =func(c int8, r int8)(int8,int8){	return 	-r,	4-c} 	// 							rho -5
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return 	-r,	4-c}} 	// 							rho -6
					} else 	if wf[2] == 1 { if wf[1] == 3 { a =func(c int8, r int8)(int8,int8){	return c-r, 4-c} 	// 					psi gamma rho alpha
									} else 	if wf[1] ==-4 { a =func(c int8, r int8)(int8,int8){	return c-r, 4-c} 	// 					psi -4
									} else 	if wf[1] ==-5 { a =func(c int8, r int8)(int8,int8){	return c-r, 4-c} 	// 					psi -5
									} else 	{  				a =func(c int8, r int8)(int8,int8){	return c-r, 4-c}}} 	// 					psi -6
	} else 	if wf[1] == 1 { 				if wf[0] == 3 { a =func(c int8, r int8)(int8,int8){	return 4-c, r} 		// 		chi gamma psi alpha
									} else 	if wf[0] ==-4 { a =func(c int8, r int8)(int8,int8){	return 4-c, r} 		// 		chi -4
									} else 	if wf[0] ==-5 { a =func(c int8, r int8)(int8,int8){	return 4-c, r} 		// 		chi -5
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return 4-c, r}} 	// 		chi -6
	} else 	if wf[2] == 1 { 				if wf[0] == 3 { a =func(c int8, r int8)(int8,int8){	return 4-c, c-r} 	// 					psi gamma rho alpha
									} else 	if wf[0] ==-4 { a =func(c int8, r int8)(int8,int8){	return 4-c, c-r} 	// 		chi -4
									} else 	if wf[0] ==-5 { a =func(c int8, r int8)(int8,int8){	return 4-c, c-r} 	// 		chi -5
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return 4-c, c-r}} 	// 		chi -6
	} else if wf[0] ==-1 { 	if wf[1] ==-2 { if wf[2] ==-3 { a =func(c int8, r int8)(int8,int8){	return 	-c, c-r} 	// 		phi -alpha psi -beta rho -gamma
									} else 	if wf[2] == 4 { a =func(c int8, r int8)(int8,int8){	return 	-c, c-r} 	// 		...
									} else 	if wf[2] == 5 { a =func(c int8, r int8)(int8,int8){	return 	-c, c-r}
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return 	-c, c-r}}
					} else 	if wf[2] ==-2 { if wf[1] ==-3 { a =func(c int8, r int8)(int8,int8){	return c-r, -c}
									} else 	if wf[1] == 4 { a =func(c int8, r int8)(int8,int8){	return c-r, -c}
									} else 	if wf[1] == 5 { a =func(c int8, r int8)(int8,int8){	return c-r, -c}
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return c-r, -c}}}
	} else if wf[0] ==-2 { 	if wf[1] ==-1 { if wf[2] ==-3 { a =func(c int8, r int8)(int8,int8){	return 	-c, 4-r}
									} else 	if wf[2] == 4 { a =func(c int8, r int8)(int8,int8){	return 	-c, 4-r}
									} else 	if wf[2] == 5 { a =func(c int8, r int8)(int8,int8){	return 	-c, 4-r}
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return 	-c, 4-r}}
					} else 	if wf[2] ==-1 { if wf[1] ==-3 { a =func(c int8, r int8)(int8,int8){	return c-r, 4-r}
									} else 	if wf[1] == 4 { a =func(c int8, r int8)(int8,int8){	return c-r, 4-r}
									} else 	if wf[1] == 5 { a =func(c int8, r int8)(int8,int8){	return c-r, 4-r}
									} else 	{  				a =func(c int8, r int8)(int8,int8){	return c-r, 4-r}}}
	} else if wf[1] ==-1 { 		 			if wf[0] ==-3 { a =func(c int8, r int8)(int8,int8){	return r-4, c}
									} else 	if wf[0] == 4 { a =func(c int8, r int8)(int8,int8){	return r-4, c}
									} else 	if wf[0] == 5 { a =func(c int8, r int8)(int8,int8){	return r-4, c}
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return r-4, c}}
	} else if wf[2] ==-1 { 		 			if wf[0] ==-3 { a =func(c int8, r int8)(int8,int8){	return r-4, c-r}
									} else 	if wf[0] == 4 { a =func(c int8, r int8)(int8,int8){	return r-4, c-r}
									} else 	if wf[0] == 5 { a =func(c int8, r int8)(int8,int8){	return r-4, c-r}
									} else 	{ 				a =func(c int8, r int8)(int8,int8){	return r-4, c-r}}}
	return a}
func (md Md) quantQp(gar Awf, bar Awf) Qp { var p Qp;
	for r, R := int8(0), int8(9); r < R; r++ {
		for c, C := int8(0), int8(9); c < C; c++ { var l Qv; var ac int8; var ar int8;
			l.c =c-4; l.r =r-4;
			ac =abs(l.c); ar =abs(l.r);
			if l.c < 0 { 		if l.r < 0 { l.lp=specLp(-1,-1)
						} else 	if l.r > 0 { l.lp=specLp( 1, 1)
			}} else if l.c > 0 {if l.r < 0 { l.lp=specLp(-1, 1)
						} else 	if l.r > 0 { l.lp=specLp( 1,-1)}}
			if ac <= ar { 	l.lp.gamma, l.lp.beta = md.gar( ac, ar);
			} else { 		l.lp.gamma, l.lp.beta = md.bar( ac, ar)}
			p[r][c] =l}} //colors
	return p}
var mDs [192]Md

//primal moods associate to pp, moods associate to lp??

type Pm struct{
	lp 					Lp
	pp 					Pp
	gwf,bwf 			Wf}
func precalcPrimalMoods(){
	pMs[5].gwf=defWf(0,1,2); 	pMs[5].bwf=defWf(0,1,2); 							pMs[5].lp =specLp( 2, 2) 								// pea
	pMs[6].gwf =pMs[5].gwf.modWf(0,2,1); pMs[6].bwf =pMs[5].bwf.modWf(0,1,2); 	pMs[6].lp =specLp( 2,-2); 								// hos
	pMs[4].gwf =pMs[5].gwf.modWf(0,1,2); pMs[4].bwf =pMs[5].bwf.modWf(2,1,0); 	pMs[4].lp =specLp(-1,-1); 	pMs[4].pp =specPp( 1,-1)	// car
	pMs[7].gwf =pMs[6].gwf.modWf(0,1,2); pMs[7].bwf =pMs[6].bwf.modWf(2,1,0); 	pMs[7].lp =specLp(-1, 1);  								// esc
	pMs[3].gwf =pMs[5].gwf.modWf(0,1,2); pMs[3].bwf =pMs[5].bwf.modWf(0,2,1); 	pMs[3].lp =specLp( 1, 1); 	 							// pla
	pMs[8].gwf =pMs[6].gwf.modWf(0,1,2); pMs[8].bwf =pMs[5].bwf.modWf(0,2,1); 	pMs[8].lp =specLp( 1,-1); 	pMs[8].pp =specPp( 1,-1)	// fig
	pMs[2].gwf =pMs[4].gwf.modWf(0,1,2); pMs[2].bwf =pMs[4].bwf.modWf(0,2,1); 	pMs[2].lp =specLp( 2, 2); 	pMs[2].pp =specPp(-1, 1)	// cbp
	pMs[9].gwf =pMs[7].gwf.modWf(0,1,2); pMs[9].bwf =pMs[7].bwf.modWf(0,2,1); 	pMs[9].lp =specLp( 2,-2); 	pMs[9].pp =specPp( 1,-1)	// ebf
	pMs[1].gwf =pMs[3].gwf.modWf(0,1,2); pMs[1].bwf =pMs[3].bwf.modWf(2,1,0); 	pMs[1].lp =specLp( 2, 2); 	pMs[1].pp =specPp(-1,-1)	// pbc
	pMs[10].gwf=pMs[8].gwf.modWf(0,1,2); pMs[10].bwf=pMs[8].bwf.modWf(2,1,0); 	pMs[10].lp=specLp( 2,-2); 	pMs[10].pp=specPp(-1,-1)	// fbe
	pMs[0].gwf =pMs[2].gwf.modWf(0,1,2); pMs[0].bwf =pMs[2].bwf.modWf(2,1,0); 	pMs[0].lp =specLp( 2, 2); 	pMs[0].pp =specPp( 1, 1) 	// cnp
	pMs[11].gwf=pMs[9].gwf.modWf(0,1,2); pMs[11].bwf=pMs[9].bwf.modWf(2,1,0); 	pMs[11].lp=specLp( 2,-2); 	pMs[11].pp=specPp( 1, 1)}	// fne
var pMs [12]Pm

type Gn struct{
	mas, Ac 			bool
	gwf, bwf 			Wf
	lp 					Lp
	pp 					Pp}
type Wf [3]int8 // Windform
func defWf( t1 int8, t2 int8, t3 int8) Wf { var wf Wf; wf[0]=t1; wf[1]=t2; wf[2]=t3; return wf}
func (f1 Wf) modWf(t1 int8, t2 int8, t3 int8) Wf { var f2 Wf
	f2[0] =f1[t1]; 	f2[1] =f1[t2]; 	f2[2] =f1[t3]; 	return f2}
type Lp struct{
	i,u,gamma,beta				int8}
func specLp(i int8, u int8) Lp { var lp Lp; lp.i=i; lp.u=u; return lp}
type Pp struct{
	i,u,gamma,beta 			int8}
func specPp( i int8, u int8) Pp { var pp Pp; pp.i=i; pp.u=u; return pp}
func quantPp(gamma int8, beta int8) Pp { var pp Pp; pp.gamma=gamma; pp.beta=beta; return pp}
func precalcGenders(){
	gNs[0].gwf=defWf( 3, 2, 1); 	gNs[0].bwf=defWf(-1,-3,-2); 															gNs[0].lp.i= 1;gNs[0].pp.i=-1;
	gNs[1].mas=true;						gNs[1].gwf=gNs[0].gwf.modWf(2,1,0); 	gNs[1].bwf=gNs[0].bwf.modWf(0,1,2); 	gNs[1].lp.i=-1;gNs[1].pp.i= 1;
						gNs[2].Ac =true;	gNs[2].gwf=gNs[1].gwf.modWf(1,0,2); 	gNs[2].bwf=gNs[1].bwf.modWf(0,1,2); 	gNs[2].lp.i= 1;gNs[2].pp.i=-1;
	gNs[3].mas=true; 	gNs[3].Ac =true;	gNs[3].gwf=gNs[1].gwf.modWf(0,1,2); 	gNs[3].bwf=gNs[1].bwf.modWf(1,0,2);	gNs[3].lp.i=-1;gNs[3].pp.i= 1}
var gNs [4]Gn



func precalcMetaphysics(){
	precalcGenders()
	precalcPrimalMoods()
	precalcMoods()

 	precalcInformationIsotopes()
	precalcFunctions()
	precalcCircuits()

	precalcSociotypes()
	precalcIntertypeRelations()

	precalcClassicalQualities()
	precalcRomanticQualities()}




//   // /// // // //////   // /// // // //////   // /// // // //////   // /// // // //////   // /// // // //////
//   // // /// //   //     // // /// //   //     // // /// //   //     // // /// //   //     // // /// //   //  
//   // //  // //   //     // //  // //   //     // //  // //   //     // //  // //   //     // //  // //   //  


func init(){

	//precalcPlanarColors()
	//precalcCubicColors()

	precalcMetaphysics()

 /*	mDs2 := mDs
 	mDs2[0].sx.mas =true
 	log.Println(mDs[0].sx.mas)
 	log.Println(mDs2[0].sx.mas)*/

	initEssences()
	initUser()

 	//log.Println( gNs)
 	//log.Println( pMs)
 	//log.Println( mDs)
}


//   // /// // // //////   // /// // // //////   // /// // // //////   // /// // // //////   // /// // // //////
//   // // /// //   //     // // /// //   //     // // /// //   //     // // /// //   //     // // /// //   //  
//   // //  // //   //     // //  // //   //     // //  // //   //     // //  // //   //     // //  // //   //  



type Color struct{
	s, v 				uint8
	h 					uint16
	hex 				string}
type ColorS []Color
func (c Color) defByHSV(h uint16,s uint8,v uint8){
						c.h=h; 	c.s=s; 	c.v=v}
func (c Color) defByHex(hex string){ c.hex=hex}
func (ca1 ColorS) spectrum( n uint8, c1 Color, c2 Color) {
	var sstep, vstep uint8;
	var sStep, vStep uint8;
	var c Color; var ca2 ColorS;
	sStep =(c2.s-c1.s)/n; 		sstep =uint8(sStep);
	vStep =(c2.v-c1.v)/n; 		vstep =uint8(vStep);
	for i:=uint8(0); i<n; i++{	c.defByHSV( c1.h, min(c1.s,c2.s)+c.s*sstep, min(c1.v,c2.v)+c.v*vstep); ca2 =append(ca2, c)}; ca1 =ca2}

func precalcPlanarColors(){
	pCs.whi.defByHSV(   1,   1,  25, 	   1,   1,  50, 	   1,   1, 100)
	pCs.red.defByHSV(   0,  25, 100, 	   0,  50,  50, 	   0, 100, 100)
	pCs.cya.defByHSV( 180,  25, 100, 	 180,  50,  50, 	 180, 100, 100)
	pCs.lil.defByHSV( 315,  25, 100, 	 315,  50,  50, 	 315, 100, 100)
	pCs.gre.defByHSV( 135,  25, 100, 	 135,  50,  50, 	 135, 100, 100)
	pCs.ora.defByHSV(  30,  25, 100, 	  30,  50,  50, 	  30, 100, 100)
	pCs.blu.defByHSV( 222,  25, 100, 	 222,  50,  50, 	 222, 100, 100)
	pCs.yel.defByHSV(  70,  25, 100, 	  70,  50,  50, 	  70, 100, 100)
	pCs.vio.defByHSV( 275,  25, 100, 	 275,  50,  50, 	 275, 100, 100)
/*	pCs.src4, pCs.src5 =pCs.simpleBlend( 	pCs.red, pCs.cya)
	pCs.slg4, pCs.slg5 =pCs.simpleBlend( 	pCs.lil, pCs.gre)
	pCs.sob4, pCs.sob5 =pCs.simpleBlend( 	pCs.ora, pCs.blu)
	pCs.syv4, pCs.syv5 =pCs.simpleBlend( 	pCs.yel, pCs.vio)
	pCs.crc4, pCs.crc5 =pCs.complexBlend( 	pCs.red, pCs.cya)
	pCs.clg4, pCs.clg5 =pCs.complexBlend( 	pCs.lil, pCs.gre)
	pCs.cob4, pCs.cob5 =pCs.complexBlend( 	pCs.ora, pCs.blu)
	pCs.cyv4, pCs.cyv5 =pCs.complexBlend( 	pCs.yel, pCs.vio)*/}
type Pc struct{
 	lo, fa, hi 			Color}
func (pc Pc) defByHSV( 	loh uint16,  los uint8,  lov uint8, 	fah uint16,  fas uint8,  fav uint8, 	hih uint16,  his uint8,  hiv uint8){
						pc.lo.h=loh;pc.lo.s=los;pc.lo.v=lov;pc.fa.h=fah;pc.fa.s=fas;pc.fa.v=fav;pc.hi.h=hih;pc.hi.s=his;pc.hi.v=hiv}
type Pcs struct{
 	whi, red, cya, lil, gre, ora, blu, yel, vio Pc
 	src, slg, sob, syv, crc, clg, cob, cyn	[][]Color}
var pCs Pcs
/*func (pcs Pcs) simpleBlend( g Pc, b Pc) ([][]Color, [][]Color) { var cm4 [][]Color; var cm5 [][]Color;
	var glf []Color;	var glh []Color;
	var blf []Color;	var blh []Color;
	glf.spectrum( 4, g.lo, g.fa);	glh.spectrum( 4, g.lo, g.hi);
	blf.spectrum( 3, b.lo, b.fa);	blh.spectrum( 3, b.lo, b.hi);
	cm4.simpleQuadrant( 4, glf, glh, blf, blh);
	glf.spectrum( 5, g.lo, g.fa);	glh.spectrum( 5, g.lo, g.hi);
	blf.spectrum( 4, b.lo, b.fa);	blh.spectrum( 4, b.lo, b.hi);
	cm5.simpleQuadrant( 5, glf, glh, blf, blh); return cm4, cm5}*/
/*func (cm1 [][]Color) simpleQuadrant( res int8, glf []Color, glh []Color, blf []Color, blh []Color) { var c Color; var cm2 [][]Color;
	for r :=int8(0); r < res; r++{
		for c :=int8(0); c < res; c++{

		}
	}}*/

/*
pattern persistence <- ability to sustain own existence
conceivable situations in which pattern (belief) evaporates, versus other cases
*/

/*func (pcs Pcs) complexBlend( g Pc, b Pc) ([][]Color, [][]Color) { var cm4 [][]Color; var cm5 [][]Color;
	var glf []Color;	var glh []Color; 	var gfh []Color;
	var blf []Color;	var blh []Color; 	var bfh []Color;
	glf.spectrum( 4, g.lo, g.fa);	glh.spectrum( 4, g.lo, g.hi); 	gfh.spectrum( 4, g.fa, g.hi);
	blf.spectrum( 3, b.lo, b.fa);	blh.spectrum( 3, b.lo, b.hi); 	bfh.spectrum( 3, b.fa, b.hi);
	cm4.complexQuadrant( glf, glh, gfh, blf, blh, bfh);
	glf.spectrum( 5, g.lo, g.fa);	glh.spectrum( 5, g.lo, g.hi); 	gfh.spectrum( 5, g.fa, g.hi);
	blf.spectrum( 4, b.lo, b.fa);	blh.spectrum( 4, b.lo, b.hi); 	bfh.spectrum( 4, b.fa, b.hi);
	cm5.complexQuadrant( glf, glh, gfh, blf, blh, bfh); return cm4, cm5}*/
/*func (cm1 [][]Color) complexQuadrant( res int8,glf []Color,glh []Color,gfh []Color,blf []Color,blh []Color,bfh []Color) {
	var c Color; var cm2 [][]Color;
	for r :=int8(0); r < res; r++{
		for c :=int8(0); c < res; c++{

		}
	}}*/





func precalcCubicColors(){
	Ccs[0].defByHex( "#ff0000")
	Ccs[1].defByHex( "#ff8888")
	Ccs[2].defByHex( "#a50000")
	Ccs[3].defByHex( "#ffa500")
	Ccs[4].defByHex( "#ffff00")
	Ccs[5].defByHex( "#a5a500")
	Ccs[6].defByHex( "#00ff00")
	Ccs[7].defByHex( "#00ffff")
	Ccs[8].defByHex( "#00a5ff")
	Ccs[9].defByHex( "#8800e1")
	Ccs[10].defByHex("#ff00ff")
	Ccs[11].defByHex("#ffa5ff")}
var Ccs [12]Color


//   ///   ///   ///   // /// //   ///   ///   ///   // /// //   ///   ///   ///   // /// //   ///   ///   ///   // /// //   ///   ///   ///   // /// //  
//   // / / //  /// /  // // ///   // / / //  /// /  // // ///   // / / //  /// /  // // ///   // / / //  /// /  // // ///   // / / //  /// /  // // ///  
//   //  /  // /////// // //  //   //  /  // /////// // //  //   //  /  // /////// // //  //   //  /  // /////// // //  //   //  /  // /////// // //  //  


func main(){
	//initIo()
	initWebSocket()	
}


//   ///   ///   ///   // /// //   ///   ///   ///   // /// //   ///   ///   ///   // /// //   ///   ///   ///   // /// //   ///   ///   ///   // /// //  
//   // / / //  /// /  // // ///   // / / //  /// /  // // ///   // / / //  /// /  // // ///   // / / //  /// /  // // ///   // / / //  /// /  // // ///  
//   //  /  // /////// // //  //   //  /  // /////// // //  //   //  /  // /////// // //  //   //  /  // /////// // //  //   //  /  // /////// // //  // 

// mental prod cant accelerate & vital acc cant accelerate?

// parent ööl pointer to own type

// can interface generalise over methods such as lon and lat for several types?


func ask(txt string) string { 
	switch txt{
		case "md": 	log.Println("md"); return "md"
	}
	return ""
}




var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" { http.Error(w, "Not found", http.StatusNotFound);	return}
	if r.Method != "GET" { http.Error(w, "Method not allowed", http.StatusMethodNotAllowed); return}
	http.ServeFile(w, r, "index1818.html")}

func initWebSocket() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { serveWs(hub, w, r) })
	err := http.ListenAndServe(*addr, nil)
	if err != nil {	log.Fatal("ListenAndServe: ", err) }}

// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {	
	clients map[*Client]bool // Registered clients.	
	broadcast chan []byte // Inbound messages from the clients.	
	register chan *Client // Register requests from the clients.
	unregister chan *Client} // Unregister requests from clients.

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),}}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register: 	h.clients[client] = true
		case client := <-h.unregister:	if _, ok := h.clients[client]; ok {
											delete(h.clients, client)
											close(client.send)}
		case message := <-h.broadcast:
			txt := string(message)
			if txt =="q"{ os.Exit(0)
			} else {
				log.Printf(txt)
				//var resp = ask(txt)
						
				for client := range h.clients {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)}}}}}}



const (
	writeWait = 10 * time.Second // Time allowed to write a message to the peer.
	pongWait = 60 * time.Second // Time allowed to read the next pong message from the peer.
	pingPeriod = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512) // Maximum message size allowed from peer.
var (
	newline = []byte{'\n'}
	space   = []byte{' '})
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,}
// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub	
	conn *websocket.Conn // The websocket connection.	
	send chan []byte} // Buffered channel of outbound messages.
// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}
// writePump pumps messages from the hub to the websocket connection.
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return}
	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()}







// MySQL

func database() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/database");

	cE( err,"Open")
	err = db.Ping(); cE( err,"Ping")

	//var initDone int
	//err = db.QueryRow("select initDone from qZ").Scan(&initDone); cE( err,"QueryRow")
	//fmt.Println( initDone)

	//stmt, err := db.Prepare("INSERT INTO qZ(initDone) VALUES(?)")
	//if err != nil { fmt.Printf("err Prepare\n")}
	//_, err := db.Exec("INSERT INTO qZ (initDone) VALUES (1)"); cE(err)
	defer db.Close()
}

/*// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // The associated file.
    Err error    // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
PathError's Error generates a string like this:

open /etc/passwx: no such file or directory*/

func cE(err error, str string) { if err != nil {
		sli := []string{}
		sli = append( sli, "Error: ")
		sli = append( sli, str)
		sli = append( sli, "\n")
		ret := strings.Join( sli,"")
		log.Printf( ret)}}
