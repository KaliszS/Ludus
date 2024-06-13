{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in with pkgs; rec {
        devShells = rec {
          goEnv = mkShell {
            name = "ludus";
            buildInputs = [
              go
            ];
            shellHook = ''
              echo "Welcome $(whoami) to the ludus shell!"
              go version
            '';
          };

          default = mkShell { inputsFrom = [ goEnv ]; };
        };
      }
    );
}
