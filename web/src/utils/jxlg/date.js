export const date = (date) => {
  const dateUTC = new Date(date)
  return dateUTC.getFullYear() +
        '-' + dateUTC.getMonth() +
        '-' + dateUTC.getDate() +
        ' ' + dateUTC.toLocaleTimeString()
}
