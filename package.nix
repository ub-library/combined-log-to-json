{buildGoModule}:
buildGoModule rec {
  name = "combined-log-to-json";
  src = builtins.path {
    path = ./.;
    name = "${name}-src";
  };
  vendorHash = null;
}
