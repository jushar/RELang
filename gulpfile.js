'use strict';
const gulp = require('gulp');
const execFileSync = require('child_process').execFileSync;
const platform = require('os').platform;
const fs = require('fs')
const join = require('path').join;

gulp.task('generate-parser', async () => {
  const args = ['-o', 'pkg/parser', '-Xexact-output-dir', '-Dlanguage=Go', 'grammar\\RELang.g4'];

  if (platform() == 'win32') {
    execFileSync('cmd.exe', ['/c', 'tools\\antlr.bat', ...args]);
  }
});

gulp.task('build', async () => {
  execFileSync('go', ['build', 'cmd/relang.go']);
});

gulp.task('run-examples', async () => {
  for (const dir of fs.readdirSync('examples')) {
    const fullPath = join('examples', dir)
    const stats = fs.statSync(fullPath);
    if (stats.isFile() && dir.endsWith('.relang')) {
      console.log('Running: ' + fullPath);

      execFileSync('go', ['run', 'cmd/relang.go', '-p', fullPath, '-o', join('examples', 'out', dir.replace(/relang$/, 'h'))]);
    }
  }
});
