import {isExternal, validTeacherId} from '@/utils/validate.js'

describe('Utils:validate', () => {
  it('validUsername', () => {
    expect(validTeacherId('admin')).toBe(true)
    expect(validTeacherId('editor')).toBe(true)
    expect(validTeacherId('xxxx')).toBe(false)
  })
  it('isExternal', () => {
    expect(isExternal('https://github.com/PanJiaChen/vue-element-admin')).toBe(true)
    expect(isExternal('http://github.com/PanJiaChen/vue-element-admin')).toBe(true)
    expect(isExternal('github.com/PanJiaChen/vue-element-admin')).toBe(false)
    expect(isExternal('/dashboard')).toBe(false)
    expect(isExternal('./dashboard')).toBe(false)
    expect(isExternal('dashboard')).toBe(false)
  })
})
