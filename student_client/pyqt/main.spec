# -*- mode: python ; coding: utf-8 -*-
# Spec file for PyInstaller
# To build project, use `pipenv run pyinstaller main.spec`, you'll find the bundle in the `dist` folder
# But before that, remember to download Python Binaries and put it in the project's root directory,
# i.e. `python-3.7.9-embed-amd64` in my case

block_cipher = None


a = Analysis(['main.py'],
             # replace this with where your dependencies lay, i.e. your output for `pipenv --venv`
             pathex=['C:\\\\Users\\\\John Mactavish\\\\.virtualenvs\\\\pyqt-Jt_vAwyy'],
             binaries=[],
             # data files copy, see documentation for details
             datas=[('index.html','.'),('static','static'),('python-3.7.9-embed-amd64','runner')],
             hiddenimports=[],
             hookspath=[],
             runtime_hooks=[],
             excludes=[],
             win_no_prefer_redirects=False,
             win_private_assemblies=False,
             cipher=block_cipher,
             noarchive=False)
pyz = PYZ(a.pure, a.zipped_data,
             cipher=block_cipher)

exe = EXE(pyz,
          a.scripts, 
          [],
          exclude_binaries=True,
          name='main',
          debug=False,
          bootloader_ignore_signals=False,
          strip=False,
          upx=True,
          console=True,
          disable_windowed_traceback=False,
          target_arch=None,
          codesign_identity=None,
          entitlements_file=None )
coll = COLLECT(exe,
               a.binaries,
               a.zipfiles,
               a.datas, 
               strip=False,
               upx=True,
               upx_exclude=[],
               name='main')
