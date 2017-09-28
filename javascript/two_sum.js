const given = [1,2,4,3,6];

function getPairWithSum(list, sum, cb){

	let map = {};
	for(let i =0; i< list.length; ++i){

		let currentValue = list[i];
		let storedVal = map[sum - currentValue];
		
		if((storedVal + currentValue) == sum){
			return cb([storedVal,currentValue]);
		}

		map[currentValue] = currentValue;
	};
}

getPairWithSum(given,8, function(res){
	console.log(res);
});

