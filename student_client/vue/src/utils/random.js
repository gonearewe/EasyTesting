export function shuffle(arr,seed) {
  let seedrandom = require('seedrandom')(seed)
  let i = arr.length, j, temp
  while(--i > 0){
    j = Math.floor(seedrandom.quick()*(i+1));
    temp = arr[j];
    arr[j] = arr[i];
    arr[i] = temp;
  }
}
