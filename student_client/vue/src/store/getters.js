const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  name: state => state.user.name,
  teacher_id: state => state.user.teacher_id,
  id: state => state.user.id,
}
export default getters
