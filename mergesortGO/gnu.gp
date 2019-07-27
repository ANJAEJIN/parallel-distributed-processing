#2018/04. gnu.gp
set term png enhanced size 640,480
set output "gnu.png"
set title "runtime, CPU#"
set xlabel "CPU#"
set ylabel "runtime"
set xrange[0:13]
set yrange[160:780]
plot "dat1.txt" t "" w l
set term aqua