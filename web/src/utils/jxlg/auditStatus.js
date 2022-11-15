export const auditStatus = (auditStatus) => {
  switch (auditStatus) {
    case 1:
      return '待审核'
    case 2:
      return '审批通过'
    case 3:
      return '审核拒绝'
    default:
      return auditStatus
  }
}
