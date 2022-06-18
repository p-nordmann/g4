/*
G4 is an open-source board game inspired by the popular game of connect-4.
Copyright (C) 2022  Pierre-Louis Nordmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package bits

// rotationLookup provides a mapping {column}->{corresponding rotated line}.
var rotationLookup = [...]bitboard{
	0,
	72057594037927936,
	281474976710656,
	72339069014638592,
	1099511627776,
	72058693549555712,
	282574488338432,
	72340168526266368,
	4294967296,
	72057598332895232,
	281479271677952,
	72339073309605888,
	1103806595072,
	72058697844523008,
	282578783305728,
	72340172821233664,
	16777216,
	72057594054705152,
	281474993487872,
	72339069031415808,
	1099528404992,
	72058693566332928,
	282574505115648,
	72340168543043584,
	4311744512,
	72057598349672448,
	281479288455168,
	72339073326383104,
	1103823372288,
	72058697861300224,
	282578800082944,
	72340172838010880,
	65536,
	72057594037993472,
	281474976776192,
	72339069014704128,
	1099511693312,
	72058693549621248,
	282574488403968,
	72340168526331904,
	4295032832,
	72057598332960768,
	281479271743488,
	72339073309671424,
	1103806660608,
	72058697844588544,
	282578783371264,
	72340172821299200,
	16842752,
	72057594054770688,
	281474993553408,
	72339069031481344,
	1099528470528,
	72058693566398464,
	282574505181184,
	72340168543109120,
	4311810048,
	72057598349737984,
	281479288520704,
	72339073326448640,
	1103823437824,
	72058697861365760,
	282578800148480,
	72340172838076416,
	256,
	72057594037928192,
	281474976710912,
	72339069014638848,
	1099511628032,
	72058693549555968,
	282574488338688,
	72340168526266624,
	4294967552,
	72057598332895488,
	281479271678208,
	72339073309606144,
	1103806595328,
	72058697844523264,
	282578783305984,
	72340172821233920,
	16777472,
	72057594054705408,
	281474993488128,
	72339069031416064,
	1099528405248,
	72058693566333184,
	282574505115904,
	72340168543043840,
	4311744768,
	72057598349672704,
	281479288455424,
	72339073326383360,
	1103823372544,
	72058697861300480,
	282578800083200,
	72340172838011136,
	65792,
	72057594037993728,
	281474976776448,
	72339069014704384,
	1099511693568,
	72058693549621504,
	282574488404224,
	72340168526332160,
	4295033088,
	72057598332961024,
	281479271743744,
	72339073309671680,
	1103806660864,
	72058697844588800,
	282578783371520,
	72340172821299456,
	16843008,
	72057594054770944,
	281474993553664,
	72339069031481600,
	1099528470784,
	72058693566398720,
	282574505181440,
	72340168543109376,
	4311810304,
	72057598349738240,
	281479288520960,
	72339073326448896,
	1103823438080,
	72058697861366016,
	282578800148736,
	72340172838076672,
	1,
	72057594037927937,
	281474976710657,
	72339069014638593,
	1099511627777,
	72058693549555713,
	282574488338433,
	72340168526266369,
	4294967297,
	72057598332895233,
	281479271677953,
	72339073309605889,
	1103806595073,
	72058697844523009,
	282578783305729,
	72340172821233665,
	16777217,
	72057594054705153,
	281474993487873,
	72339069031415809,
	1099528404993,
	72058693566332929,
	282574505115649,
	72340168543043585,
	4311744513,
	72057598349672449,
	281479288455169,
	72339073326383105,
	1103823372289,
	72058697861300225,
	282578800082945,
	72340172838010881,
	65537,
	72057594037993473,
	281474976776193,
	72339069014704129,
	1099511693313,
	72058693549621249,
	282574488403969,
	72340168526331905,
	4295032833,
	72057598332960769,
	281479271743489,
	72339073309671425,
	1103806660609,
	72058697844588545,
	282578783371265,
	72340172821299201,
	16842753,
	72057594054770689,
	281474993553409,
	72339069031481345,
	1099528470529,
	72058693566398465,
	282574505181185,
	72340168543109121,
	4311810049,
	72057598349737985,
	281479288520705,
	72339073326448641,
	1103823437825,
	72058697861365761,
	282578800148481,
	72340172838076417,
	257,
	72057594037928193,
	281474976710913,
	72339069014638849,
	1099511628033,
	72058693549555969,
	282574488338689,
	72340168526266625,
	4294967553,
	72057598332895489,
	281479271678209,
	72339073309606145,
	1103806595329,
	72058697844523265,
	282578783305985,
	72340172821233921,
	16777473,
	72057594054705409,
	281474993488129,
	72339069031416065,
	1099528405249,
	72058693566333185,
	282574505115905,
	72340168543043841,
	4311744769,
	72057598349672705,
	281479288455425,
	72339073326383361,
	1103823372545,
	72058697861300481,
	282578800083201,
	72340172838011137,
	65793,
	72057594037993729,
	281474976776449,
	72339069014704385,
	1099511693569,
	72058693549621505,
	282574488404225,
	72340168526332161,
	4295033089,
	72057598332961025,
	281479271743745,
	72339073309671681,
	1103806660865,
	72058697844588801,
	282578783371521,
	72340172821299457,
	16843009,
	72057594054770945,
	281474993553665,
	72339069031481601,
	1099528470785,
	72058693566398721,
	282574505181441,
	72340168543109377,
	4311810305,
	72057598349738241,
	281479288520961,
	72339073326448897,
	1103823438081,
	72058697861366017,
	282578800148737,
	72340172838076673,
}
