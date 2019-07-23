#2018/04. gnu.gp
set term png enhanced size 640,480
set output "gnu.png"
set title "runtime, CPU#"
set xlabel "CPU#"
set ylabel "runtime"
set xrange[0:12]
set yrange[40:210]
plot "dat1.txt" t "" w p
set term aqua