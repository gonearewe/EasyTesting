const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.user.token,
  name: state => state.user.name,
  teacher_id: state => state.user.teacher_id,
  id: state => state.user.id,
  roles: state => state.user.roles,
  permission_routes: state => state.permission.routes,
}
export default getters
