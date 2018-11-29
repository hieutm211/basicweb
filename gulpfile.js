
var gulp = require("gulp")
var shell = require("gulp-shell")

gulp.task("install-binary", shell.task([
	"go install"
]))

gulp.task("restart-supervisor", ["install-binary"], shell.task([
	"sudo supervisorctl restart basicweb"
]))

gulp.task("watch", function() {
	gulp.watch("*", ["install-binary", "restart-supervisor"])
})

gulp.task("default", ["watch"])
