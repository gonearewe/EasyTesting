const getters = {
  sidebar: state => state.app.sidebar,
  token: state => state.user.token,
  name: state => state.user.name,
  student_id: state => state.user.student_id,
  class_id: state => state.user.class_id,
  exam_session_id: state => state.user.exam_session_id,
  exam_deadline: state => state.user.exam_deadline,
}
export default getters
