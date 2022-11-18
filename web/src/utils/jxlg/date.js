export const date = (date) => {
  const dateUTC = new Date(date)
  return dateUTC.getFullYear() +
        '-' + (dateUTC.getMonth() + 1) +
        '-' + dateUTC.getDate() +
        ' ' + dateUTC.toLocaleTimeString('chinese', { hour12: false })
}
