gcloud app versions list

echo ""
echo "Which version delete?"

read INPUT
gcloud app versions delete ${INPUT}
