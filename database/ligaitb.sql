-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 12, 2017 at 03:26 PM
-- Server version: 10.1.28-MariaDB
-- PHP Version: 7.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ligaitb`
--

-- --------------------------------------------------------

--
-- Table structure for table `klasemen`
--

CREATE TABLE `klasemen` (
  `rank` int(2) NOT NULL,
  `himpunan` varchar(12) NOT NULL,
  `tanding` int(2) NOT NULL,
  `menang` int(2) NOT NULL,
  `seri` int(2) NOT NULL,
  `kalah` int(2) NOT NULL,
  `gm` int(11) NOT NULL,
  `gk` int(11) NOT NULL,
  `sg` int(11) NOT NULL,
  `poin` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `klasemen`
--

INSERT INTO `klasemen` (`rank`, `himpunan`, `tanding`, `menang`, `seri`, `kalah`, `gm`, `gk`, `sg`, `poin`) VALUES
(13, 'AMISCA', 34, 9, 14, 11, 39, 34, 5, 41),
(12, 'HIMATG TERRA', 34, 12, 6, 16, 46, 58, -12, 42),
(17, 'HIMATIKA', 34, 7, 6, 21, 41, 62, -21, 27),
(8, 'HME', 34, 15, 7, 12, 50, 39, 11, 52),
(3, 'HMFT', 34, 19, 8, 7, 67, 38, 29, 65),
(7, 'HMIF', 34, 15, 8, 11, 48, 44, 4, 53),
(4, 'HMM', 34, 17, 10, 7, 46, 24, 22, 61),
(18, 'HMP', 34, 2, 4, 28, 26, 104, -78, 10),
(10, 'HMS', 34, 13, 4, 17, 49, 74, -25, 43),
(1, 'HMT', 34, 22, 2, 10, 61, 40, 21, 68),
(14, 'HMTG GEA', 34, 12, 4, 18, 49, 55, -6, 40),
(15, 'HMTM PATRA', 34, 10, 7, 17, 35, 45, -10, 37),
(9, 'IMG', 34, 13, 10, 11, 43, 44, -1, 49),
(5, 'IMMG', 34, 17, 9, 8, 58, 44, 14, 60),
(16, 'KMM', 34, 9, 8, 17, 34, 52, -18, 35),
(2, 'KMPN', 34, 21, 5, 8, 76, 38, 38, 68),
(11, 'MTI', 34, 11, 9, 14, 50, 50, 0, 42),
(6, 'MTM', 34, 17, 9, 8, 64, 37, 27, 60);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `klasemen`
--
ALTER TABLE `klasemen`
  ADD PRIMARY KEY (`himpunan`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
