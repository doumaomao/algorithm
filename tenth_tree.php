<?php
outputDict(1000,1);
function outputDict($max_num,$k) {
	for($i=0; $i<10&&$i+$k<$max_num; $i++){
		if($i==9 && $k==1){
			break;
		}
		echo $i+$k."\r\n";
		outputDict($max_num, ($i+$k)*10);
	}
}
