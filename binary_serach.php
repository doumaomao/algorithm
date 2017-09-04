<?php
	// init value
	$intFindValue = $argv[1];
	$arrData =  array(1, 2, 3, 4, 5);
	$left = 0;
	$right = count($arrData) - 1;
	
	// find value
	while($left <= $right) {
		$middle = $left + (($right-$left) >> 1);
		if($arrData[$middle] > $intFindValue) {
			$right = $middle - 1;
		}elseif($arrData[$middle] < $intFindValue){
			$left = $middle + 1;
		}else{
			return $middle;
		}
	}
	return false;
?>
