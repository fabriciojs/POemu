echo "Compiling...";

go build po.go;

echo "Done.";
echo "Injectin...";

echo "palmeiras...";
./bin/poemu --query="palmeiras" --cod_busca=67 --media=twitter >> output.log 2>&1 &
./bin/poemu --query="palmeiras" --cod_busca=67 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "corinthians...";
./bin/poemu --query="corinthians" --cod_busca=68 --media=twitter >> output.log 2>&1 &
./bin/poemu --query="corinthians" --cod_busca=68 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Federal...";
./bin/poemu --query="Federal" --cod_busca=69 --media=twitter >> output.log 2>&1 &
./bin/poemu --query="Federal" --cod_busca=69 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Brasil...";
./bin/poemu --query="Brasil" --cod_busca=70 --media=twitter >> output.log 2>&1 &
./bin/poemu --query="Brasil" --cod_busca=70 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Papai Noel...";
./bin/poemu --query="Papai Noel" --cod_busca=71 --media=twitter >> output.log 2>&1 &
./bin/poemu --query="Papai Noel" --cod_busca=71 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Natal...";
./bin/poemu --query="Natal" --cod_busca=72 --media=twitter >> output.log 2>&1 &
./bin/poemu --query="Natal" --cod_busca=72 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "bom dia...";
./bin/poemu --query="bom dia" --cod_busca=73 --media=twitter >> output.log 2>&1 &
./bin/poemu --query="bom dia" --cod_busca=73 --media=facebook >> output.log 2>&1 &

echo "Done.";