echo "Compiling...";

go build po.go;

echo "Done.";
echo "Injectin...";

echo "palmeiras...";
./po --query="palmeiras" --cod_busca=67 --media=twitter >> output.log 2>&1 &
./po --query="palmeiras" --cod_busca=67 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "corinthians...";
./po --query="corinthians" --cod_busca=68 --media=twitter >> output.log 2>&1 &
./po --query="corinthians" --cod_busca=68 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Federal...";
./po --query="Federal" --cod_busca=69 --media=twitter >> output.log 2>&1 &
./po --query="Federal" --cod_busca=69 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Brasil...";
./po --query="Brasil" --cod_busca=70 --media=twitter >> output.log 2>&1 &
./po --query="Brasil" --cod_busca=70 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Papai Noel...";
./po --query="Papai Noel" --cod_busca=71 --media=twitter >> output.log 2>&1 &
./po --query="Papai Noel" --cod_busca=71 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "Natal...";
./po --query="Natal" --cod_busca=72 --media=twitter >> output.log 2>&1 &
./po --query="Natal" --cod_busca=72 --media=facebook >> output.log 2>&1 &
sleep 5;
echo "bom dia...";
./po --query="bom dia" --cod_busca=73 --media=twitter >> output.log 2>&1 &
./po --query="bom dia" --cod_busca=73 --media=facebook >> output.log 2>&1 &

echo "Done.";