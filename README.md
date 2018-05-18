# seguranca-formacaoPalavras

<h3>Questão 3 (Método formWords)</h3>
<p><b>Enunciado:</b></p>
<p>Construa um algorítmo que produza todas as palavra de 4 letras formadas por uma consoante seguinda de uma vogal.</p>
<br/>
<p><b>Solução</b></p>
<p>Para solucionar esta questão foi implementando um método que possui dois arrays de string, um para as cosoantes e um para as voagias. A partir daí utiliza-se um contador para verificar quais são todas as palavras possíveis. Para encotrar as palvras foi utilizado 4 estruturas do tipo for na seguinte ordem:</p>

`consoante - vogal - consoante - vogal`
<br/><br/>

<h3>Questão 4 (Método formWordsWithoutRepetition)</h3>
<p><b>Enunciado:</b></p>
<p>Construa um algorítmo que produza todas as palavra de 6 letras formadas por uma consoante seguinda de uma vogal, mas que cada sílaba (consoante+vogal) não seja repetida.</p>
<br/>
<p><b>Solução</b></p>
<p>Para solucionar esta questão foi implementando um método que possui dois arrays de string, um para as cosoantes e um para as voagias. A partir daí utiliza-se um contador para verificar quais são todas as palavras possíveis. Para encotrar as palvras foi utilizado 6 estruturas do tipo for na seguinte ordem:</p>

`consoante - vogal - consoante - vogal - consoante - vogal`

<p>A cada sílaba formada (consoante + vogal) cria-se uma variavel para sílaba, quando for criada a segunda sílaba verifica se é igual com a primeira, se for dirente cria-se a terceira sílaba e verifica-se se ela é diferente de todas as outas, caso seja imprime a palavra mas caso em qualquer uma dessas condições as sílabas sejam iguais a palavra é ignorada. </p>
<br/><br/>

<h2>Uso:</h2>

<p>Utilize a Opção -q para definir qual a questão que será executada</p>

<p>Exemplo de uso para questão 3</p>
`./main -q 3`

<p>Exemplo de uso para questão 4</p>
`./main -q 4`
