package main

var juliaDockerFile string =`FROM gitpod/workspace-full

USER gitpod

# Install Julia
RUN sudo apt-get update \
	&& sudo apt-get install -y \
		build-essential \
		libatomic1 \
		python \
		gfortran \
		perl \
		wget \
		m4 \
		cmake \
		pkg-config \
		julia \
	&& sudo rm -rf /var/lib/apt/lists/*

# Give control back to Gitpod Layer
USER root`

var juliaYaml string = `image:
  file: .gitpod.Dockerfile

tasks:
- command: julia --version

vscode:
  extensions:
    - julialang.language-julia@0.12.3:lgRyBd8rjwUpMGG0C5GAig==
`

var nimDockerFile string = `FROM gitpod/workspace-full

USER gitpod

RUN sudo apt-get update \
	&& sudo apt-get install -y \
		nim`

var nimYaml string = `image:
  file: .gitpod.Dockerfile

tasks:
  - command: nimc --version

vscode:
  extensions:
	- kosz78.nim@0.6.3:w7n1wKOFVkz9yIqgRYT7lQ==`

var hyDockerfile string = `FROM gitpod/workspace-full

USER gitpod

RUN pip3 install hy --user
`

var hyYaml string = `image:
  file: .gitpod.Dockerfile

vscode:
  extensions:
	- xuqinghan.vscode-hy@0.0.4:Utf282betrZISZjOJLTZlg==`
var clojureDockerfile string = `FROM gitpod/workspace-full

USER gitpod

# Install Clojure
RUN curl -O https://download.clojure.org/install/linux-install-1.10.1.492.sh \
    && chmod +x linux-install-1.10.1.492.sh  \
    && sudo ./linux-install-1.10.1.492.sh

# Give access back to gitpod image builder
USER root
`
var clojureYaml string = `image:
  file: .gitpod.Dockerfile

vscode:
  extensions:
    - avli.clojure@0.11.1:LAV1SbBlP0gU7J8kduhQvQ==`