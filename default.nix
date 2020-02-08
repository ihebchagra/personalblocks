{ buildGoModule, fetchFromGitHub }:

buildGoModule {
  name = "personalblocks";
  src = fetchFromGitHub {
    owner = "ihebchagra";
    repo = "personalblocks";
    rev = "1276fbe2bebecef8847535e638e5484371dd26ca";
    sha256 = "1ncrsz2aci9aznadyrwbbxaplf85x9d2p5l2c26ja49mh1lpkr17";
  };
  modSha256 = "0sjjj9z1dhilhpc8pq4154czrb79z9cm044jvn75kxcjv6v5l2m5";
  subPackages = [ "." ];
}

