# CMAKE generated file: DO NOT EDIT!
# Generated by "MinGW Makefiles" Generator, CMake Version 3.16

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

SHELL = cmd.exe

# The CMake executable.
CMAKE_COMMAND = D:\Prosoftware\sdks\cmake\bin\cmake.exe

# The command to remove a file.
RM = D:\Prosoftware\sdks\cmake\bin\cmake.exe -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = D:\workspace\go\src\project1\document_server\libc\sm4

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = D:\workspace\go\src\project1\document_server\libc\sm4\build

# Include any dependencies generated for this target.
include CMakeFiles/sm4.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/sm4.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/sm4.dir/flags.make

CMakeFiles/sm4.dir/sm4.c.obj: CMakeFiles/sm4.dir/flags.make
CMakeFiles/sm4.dir/sm4.c.obj: CMakeFiles/sm4.dir/includes_C.rsp
CMakeFiles/sm4.dir/sm4.c.obj: ../sm4.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=D:\workspace\go\src\project1\document_server\libc\sm4\build\CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/sm4.dir/sm4.c.obj"
	D:\Prosoftware\sdks\gcc_home\MinGW\bin\gcc.exe $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles\sm4.dir\sm4.c.obj   -c D:\workspace\go\src\project1\document_server\libc\sm4\sm4.c

CMakeFiles/sm4.dir/sm4.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/sm4.dir/sm4.c.i"
	D:\Prosoftware\sdks\gcc_home\MinGW\bin\gcc.exe $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E D:\workspace\go\src\project1\document_server\libc\sm4\sm4.c > CMakeFiles\sm4.dir\sm4.c.i

CMakeFiles/sm4.dir/sm4.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/sm4.dir/sm4.c.s"
	D:\Prosoftware\sdks\gcc_home\MinGW\bin\gcc.exe $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S D:\workspace\go\src\project1\document_server\libc\sm4\sm4.c -o CMakeFiles\sm4.dir\sm4.c.s

# Object files for target sm4
sm4_OBJECTS = \
"CMakeFiles/sm4.dir/sm4.c.obj"

# External object files for target sm4
sm4_EXTERNAL_OBJECTS =

libsm4.dll: CMakeFiles/sm4.dir/sm4.c.obj
libsm4.dll: CMakeFiles/sm4.dir/build.make
libsm4.dll: CMakeFiles/sm4.dir/linklibs.rsp
libsm4.dll: CMakeFiles/sm4.dir/objects1.rsp
libsm4.dll: CMakeFiles/sm4.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=D:\workspace\go\src\project1\document_server\libc\sm4\build\CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking C shared library libsm4.dll"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles\sm4.dir\link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/sm4.dir/build: libsm4.dll

.PHONY : CMakeFiles/sm4.dir/build

CMakeFiles/sm4.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles\sm4.dir\cmake_clean.cmake
.PHONY : CMakeFiles/sm4.dir/clean

CMakeFiles/sm4.dir/depend:
	$(CMAKE_COMMAND) -E cmake_depends "MinGW Makefiles" D:\workspace\go\src\project1\document_server\libc\sm4 D:\workspace\go\src\project1\document_server\libc\sm4 D:\workspace\go\src\project1\document_server\libc\sm4\build D:\workspace\go\src\project1\document_server\libc\sm4\build D:\workspace\go\src\project1\document_server\libc\sm4\build\CMakeFiles\sm4.dir\DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/sm4.dir/depend

