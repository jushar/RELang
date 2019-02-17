'use strict';
const gulp = require('gulp');
const execFileSync = require('child_process').execFileSync;
const platform = require('os').platform;

gulp.task('generate-parser', async () => {
    const args = ['-o', 'pkg/parser', '-Xexact-output-dir', '-Dlanguage=Go', 'grammar\\RELang.g4'];

    if (platform() == 'win32') {
        execFileSync('cmd.exe', ['/c', 'tools\\antlr.bat', ...args]);
    }
});
