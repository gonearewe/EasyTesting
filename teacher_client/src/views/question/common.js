export function getDifficultyColor(difficulty) {
  if (difficulty.toFixed(1) >= 0.6) {
    return 'danger'
  } else if (difficulty.toFixed(1) >= 0.3) {
    return 'primary'
  } else {
    return 'success'
  }
}

