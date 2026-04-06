export function getRemainingTimeColor(remainingTimeInPercent: number) {
  if (remainingTimeInPercent < 0.1) {
    return "css-class-remaining-time-danger-color";
  } else if (remainingTimeInPercent < 0.25) {
    return "css-class-remaining-time-warning-color";
  } else if (remainingTimeInPercent < 0.5) {
    return "css-class-remaining-time-pending-color";
  }

  return ""
}
