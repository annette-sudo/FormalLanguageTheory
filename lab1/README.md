# Лабораторная работа №1
Дана SRS $\tau$:

1. $$abbaba \rightarrow aabbbaa$$
2. $$aaa \rightarrow abab$$
3. $$abba \rightarrow baaab$$
4. $$bbbb \rightarrow ba$$

## Завершимость 
Данная SRS незавершима. Пример: 

$$
w_1 abbaa w_2
$$
Применим правило 3:
$$
w_1 baaaba w_2
$$
Применим правило 2:
$$
w_1 bababba w_2
$$
Применим правило 3:
$$
w_1 babbaaab w_2
$$
Сократим:
$$
w'_1 abbaa w'_2
$$

Где $w'_1 > w_1$ и $w'_2 > w_2$. Мы получили бесконечный цикл.

## Классы эквивалентности
### Конечность классов эквивалентности по НФ
Заметим, что если слева стоит $b$,  то вся строка приводитя к виду $b^n$. Следовательно, нужно рассмотреть два набора:
1. набор, где невозможно привести $b$ влево;
2. последовательности $b^n$.

Рассмотрим первый набор, где $b$ нельзя привести влево. 
1. $a$
2. $aa, ab$
3. $aab, aba, abb$ ($aaa$ рассматривается ниже)
4. $aaba, aabb, abbb$
5. $aabbb$

К строкам $$a, aa, ab, aab, abb, aabb, abbb, aabbb$$ не применимо ни одно правило переписывания. Теперь рассмотрим $$aba \leftrightarrow abbbb$$. Как мы видим, никакие правила больше нельзя применить, $b$ нельзя привести влево. Аналогично, $$aaba \leftrightarrow aabbbb$$. 

Слова длины 4, где $b$ не стоит первой, исключая написанные выше, не подходят, так как:

$$
abaa \leftrightarrow abbbba \leftrightarrow abbbbbbb \leftrightarrow abbabb \leftrightarrow baaabbb \\
$$
$$ 
abba \leftrightarrow baaab \\
$$
$$
abab \leftrightarrow aaa \leftrightarrow baaa  \\
$$
$$
aaaa \leftrightarrow aaa \leftrightarrow baaa \\
$$
$$
aaab \leftrightarrow  baaab \\
$$

Таких слов 5 штук.

Слова длины 5, где $b$ не стоит первой, исключая $aabbb$ и строки, где префиксом являются слова длины 4, рассмотренные выше (таких слов, очевидно, 10), не подходят так как:

$$
aabaa \leftrightarrow aabbbba \leftrightarrow aab^7 \leftrightarrow aabbabb \leftrightarrow abaaabbb (получили префикс abaa, рассмотренный выше) \\
$$ 
$$
aabab \leftrightarrow aaaa \leftrightarrow baaa \\
$$
$$
aabba \leftrightarrow abaaab \leftrightarrow baaabbbab \\
$$
$$
abbba \leftrightarrow ab^6 \leftrightarrow abbaa \leftrightarrow  baaaba \\
$$
$$
abbbb \leftrightarrow aba  \\
$$

Проверим, что все такие слова длины 5 рассмотрели. Слов длины 5, где $b$ не стоит первой, всего $2^4 = 16$. Так и получается, 10 (слова с "неудачными" префиксами длины 4) + 5 (слова, рассмотренные выше) + 1 ($aabbb$). 

Теперь рассмотрим слова длины 6, где $b$ не стоит первой. Их всего $2^5 = 32$. Из предыдущих рассуждений получается, что всего 15 "неудачных" префиксов длины 5. Таким образом, можно не рассматривать 30 слов. Остаются $aabbbb$ и $aabbba$.

$$
aabbbb \leftrightarrow aaba  \\
$$
$$
aabbba \leftrightarrow aab^6 \leftrightarrow aabbab \leftrightarrow abaaabb \\
$$

В первом случае слово длины 4, а во втором "неудачный" префикс $abaa$. Дальше можно не рассматривать, так как слов длины 7, где $b$ не стоит первой, всего 64. У нас уже имеется 32 "неудачных" префикса длины 6, то есть всего они дают 72 слова. И так далее.

Теперь мы получили первый набор:

$$a, aa, ab, aab, aba, abb, aaba, aabb, abbb, aabbb$$

Следует обратить внимание на $aaa$:

$$aaa \leftrightarrow abab \leftrightarrow ab^5 \leftrightarrow abba \leftrightarrow baaab \leftrightarrow bababb \leftrightarrow b^{10} \leftrightarrow b^6ba \leftrightarrow b^3baa \leftrightarrow baaa$$

Таким образом, мы получаем 

$$aaa \leftrightarrow baaa$$

$$aaa \leftrightarrow b^{10}$$

Аналогично, 

$$aaaa \leftrightarrow ababa \leftrightarrow ab^5a \leftrightarrow abbaa \leftrightarrow baaaba \leftrightarrow bababba \leftrightarrow b^{13} \leftrightarrow b^3aaa \leftrightarrow aaa$$

То есть, 

$$aaa \leftrightarrow aaaa$$

Рассмотрим второй набор:

1. $b$
2. $bb$
3. $bbb$
4. $$bbbb \leftrightarrow ba$$
5. $$bbbbb \leftrightarrow bab$$
6. $$bbbbbb \leftrightarrow babb$$
7. $$b^7 \leftrightarrow b^4a \leftrightarrow baa $$
8. $$b^8 \leftrightarrow baab $$
9. $$b^9 \leftrightarrow b^4b^4b \leftrightarrow babab \leftrightarrow baaa \leftrightarrow aaa $$
10. $$b^{10} \leftrightarrow baaa \leftrightarrow aaa $$

При $$n \geq 9$$  $$b^n \leftrightarrow  aaa$$.

Мы получаем классы эквивалентности: $$\epsilon, a, aa, ab, aab, aba, abb, aaba, aabb, abbb, aabbb, b, bb, bbb, ba, bab, babb, baa, baab, aaa$$.


### Построение ДКА
Построим ДКА по данным классам эквивалентности.
![ДКА](img/automa.svg)

### Построение минимальной системы переписывания
Воспользовавашись опытом прошлых цивилизаций, а точнее работой ["Построение трансформационного моноида по автомату (А.Иванов)](https://github.com/UsefulTornado/Formal-Languages), мы получили минимальную систему переписывания:

$$
bba \rightarrow bab \\
$$
$$
aaaa \rightarrow aaa \\
$$
$$
aaab \rightarrow aaa \\
$$
$$
abaa \rightarrow aaa \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaa \rightarrow aaa \\
$$
$$
baba \rightarrow baab \\
$$
$$
bbbb \rightarrow ba \\
$$
$$
baaba \rightarrow aaa \\
$$
$$
baabb \rightarrow aaa \\
$$
$$
babbb \rightarrow baa \\
$$


## Локальная конфлюэнтность
### Переупорядочивание правил
Переориентируем правила, чтобы задать фундированный порядок.

$$
aabbbaa \rightarrow abbaba \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaab \rightarrow abba \\ 
$$
$$
bbbb \rightarrow ba \\
$$

Теперь SRS завершима, так как каждое правило уменьшает длину слова.

Фундированный порядок: shortlex (армейский), где $$a < b$$. 
### Проверка конфлюэнтности
Рассмотрим слово $bbbbb$:

$$bbbbb \rightarrow bab$$

$$bbbbb \rightarrow bba$$

Как мы видим, SRS не конфлюэнтна.

### Пополняемость по Кнуту-Бендикса
Доведем сиситему до конфлюэнтности. 

1. Из предыдущего пункта получаем:

$$bba \rightarrow bab$$

2. Проверим слово $bbbba$:

$$bbbba \rightarrow baa$$

$$bbbba \rightarrow bbbab \rightarrow bbabb \rightarrow babbb$$

Получаем: $$babbb \rightarrow baa$$

3. Проверим слово $babbbb$:

$$babbbb \rightarrow baab$$

$$babbbb \rightarrow baba$$

Получаем: $$baba \rightarrow baab$$

4. Проверим слово $babbba$:

$$babbba \rightarrow baaa$$

$$babbba \rightarrow babbab \rightarrow bababb \rightarrow baaab \rightarrow abba \rightarrow abab \rightarrow aaa$$

Получаем: $$baaa \rightarrow aaa$$

5. Проверим слово $baaab$:

$$baaab \rightarrow aaab$$

$$baaab \rightarrow abba \rightarrow abab \rightarrow aaa$$

Получаем: $$aaab \rightarrow aaa$$

6. Проверим слово $babab$: 

$$babab \rightarrow baabb$$

$$babab \rightarrow baaaa \rightarrow aaa$$

Получаем: $$baabb \rightarrow aaa$$

7. Проверим слово $abbbbba$: 

$$abbbbba \rightarrow ababa \rightarrow aaaa$$

$$abbbbba \rightarrow abbbbab \rightarrow abbbabb \rightarrow abbabbb \rightarrow ababbbb \rightarrow aaabbbb \rightarrow aaa$$

Получаем: $$aaaa \rightarrow aaa$$

7. Ещё раз рассмотрим слово $abbbbba$: 

$$abbbbba \rightarrow ababa \rightarrow baaaba$$

$$abbbbba \rightarrow abbbbab \rightarrow abbbabb \rightarrow abbabbb \rightarrow ababbbb \rightarrow aaabbb \rightarrow aaa$$

Получаем: $$baaaba \rightarrow aaa$$

8. Проверим слово $abbbba$: 

$$abbbba \rightarrow abaa$$

$$abbbba \rightarrow abbbab \rightarrow abbabb \rightarrow ababbb \rightarrow aaabb \rightarrow aaa$$

Получаем: $$abaa \rightarrow aaa$$


Пополненая система:

$$
aabbbaa \rightarrow abbaba \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaab \rightarrow abba \\ 
$$
$$
bbbb \rightarrow ba \\
$$
$$
bba \rightarrow bab \\
$$
$$
aaaa \rightarrow aaa \\
$$
$$
aaab \rightarrow aaa \\
$$
$$
abaa \rightarrow aaa \\
$$
$$
baaa \rightarrow aaa \\
$$
$$
baba \rightarrow baab \\
$$
$$
baaba \rightarrow aaa \\
$$
$$
baabb \rightarrow aaa \\
$$
$$
babbb \rightarrow baa \\
$$


Заметим, что при пополнение по Кнуту-Бендиксу мы получаем те же правила, что были в минимальной системе. Это объясняется тем, что из конечного числа классов, следует конечное множество переходов между ними.


## Инварианты 
Исходная SRS $\tau$ (с переориентированными правилами):

$$
aabbbaa \rightarrow abbaba \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaab \rightarrow abba \\ 
$$
$$
bbbb \rightarrow ba \\
$$

Минимальная SRS $\tau'$:

$$
bba \rightarrow bab \\
$$
$$
aaaa \rightarrow aaa \\
$$
$$
aaab \rightarrow aaa \\
$$
$$
abaa \rightarrow aaa \\
$$
$$
abab \rightarrow aaa \\
$$
$$
baaa \rightarrow aaa \\
$$
$$
baba \rightarrow baab \\
$$
$$
bbbb \rightarrow ba \\
$$
$$
baaba \rightarrow aaa \\
$$
$$
baabb \rightarrow aaa \\
$$
$$
babbb \rightarrow baa \\
$$

Инварианты:
1. Буква $a$ сохраняется.
2. Количество подстрок $bbb$ сокращается.