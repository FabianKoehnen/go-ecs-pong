{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";
    devenv.url = "github:cachix/devenv";
    devenv.inputs.nixpkgs.follows = "nixpkgs";
  };

  nixConfig = {
    extra-trusted-public-keys = "devenv.cachix.org-1:w1cLUi8dv3hnoSPGAuibQv+f9TZLr6cv/Hm9XgU50cw=";
    extra-substituters = "https://devenv.cachix.org";
  };

  outputs = { self, nixpkgs, devenv, systems, ... } @ inputs:
    let
      forEachSystem = nixpkgs.lib.genAttrs (import systems);
    in
    {
      devShells = let 
        pkgs = import nixpkgs {
          system = "x86_64-linux";
          config.allowUnfree = true;
        };
      in{
        x86_64-linux.default = pkgs.mkShell{
          packages = [
            pkgs.pkg-config
            pkgs.vips
            
            pkgs.go

            # Required by vscode-go
            pkgs.delve

            # vscode-go expects all tool compiled with the same used go version, see: https://github.com/golang/vscode-go/blob/72249dc940e5b6ec97b08e6690a5f042644e2bb5/src/goInstallTools.ts#L721
            pkgs.gotools
            pkgs.gomodifytags
            pkgs.impl
            pkgs.go-tools
            pkgs.gopls
            pkgs.gotest


            # pkgs.xorg.libX11.dev
            # pkgs.xorg.libXrandr.dev
            # pkgs.xorg.libXxf86vm.dev
            # pkgs.xorg.libXi.dev
            # pkgs.xorg.libXcursor.dev
            # pkgs.xorg.libXinerama.dev
            # pkgs.libGL.dev
            # pkgs.libGLU.dev
            # pkgs.libglvnd
            # pkgs.libgcc


            pkgs.libGL
            pkgs.xorg.libX11
            pkgs.xorg.libXrandr
            pkgs.xorg.libXcursor
            pkgs.xorg.libXinerama
            pkgs.xorg.libXi
            pkgs.xorg.libXxf86vm

          
          ];

          env.GOROOT = pkgs.go + "/share/go/";

          hardeningDisable = [ "fortify" ];

          shellHook = ''
            export LD_LIBRARY_PATH=${pkgs.wayland}/lib:${pkgs.lib.getLib pkgs.libGL}/lib:${pkgs.lib.getLib pkgs.libGL}/lib:$LD_LIBRARY_PATH
          '';

          # enterShell = ''
          #   export PATH=$GOPATH/bin:$PATH
          # '';
        };
      };
    };
}
