-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 01, 2022 at 04:30 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 7.3.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_quiz1_pbp`
--

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `id` int(11) NOT NULL,
  `idWallet` int(11) NOT NULL,
  `datetime` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `amount` bigint(20) NOT NULL,
  `description` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction`
--

INSERT INTO `transaction` (`id`, `idWallet`, `datetime`, `amount`, `description`) VALUES
(1, 1, '2022-03-01 01:19:30', 200000, 'bonus harian'),
(2, 1, '2022-03-01 01:19:35', 2000000, 'gaji bulanan'),
(3, 1, '2022-03-01 01:19:41', 100000, 'uang lembur harian'),
(4, 1, '2022-03-01 01:21:20', 100000, 'uang jajan harian'),
(5, 2, '2022-03-01 01:21:22', 10, 'bonus harian'),
(6, 2, '2022-03-01 01:21:25', 200, 'bonus harian'),
(7, 2, '2022-03-01 01:21:27', 20000, 'gaji bulanan'),
(8, 2, '2022-03-01 01:21:30', 1000, 'uang lembur harian'),
(9, 3, '2022-03-01 01:21:32', 15, 'bonus harian'),
(10, 3, '2022-03-01 01:21:35', 250, 'bonus harian'),
(11, 3, '2022-03-01 01:21:37', 25000, 'gaji bulanan'),
(12, 3, '2022-03-01 01:21:39', 1500, 'uang lembur harian'),
(13, 4, '2022-03-01 01:21:44', 25, 'bonus harian'),
(14, 4, '2022-03-01 01:21:46', 350, 'bonus harian'),
(15, 4, '2022-03-01 01:21:49', 35000, 'gaji bulanan'),
(16, 4, '2022-03-01 01:21:51', 2500, 'uang lembur harian'),
(34, 1, '2022-03-01 02:30:30', 15000, 'makan siang'),
(36, 1, '2022-03-01 03:15:32', 0, 'makan siang'),
(37, 1, '2022-03-01 03:15:40', -20000, 'makan siang');

-- --------------------------------------------------------

--
-- Table structure for table `wallet`
--

CREATE TABLE `wallet` (
  `id` int(11) NOT NULL,
  `currency` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `disableUser` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `wallet`
--

INSERT INTO `wallet` (`id`, `currency`, `username`, `password`, `disableUser`) VALUES
(1, 'IDR', 'jasti', 'jas', 1),
(2, 'SGD', 'josh', 'josh1', 0),
(3, 'AUD', 'Jack', 'jack1', 1),
(4, 'USD', 'joy', 'joy1', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idWallet` (`idWallet`);

--
-- Indexes for table `wallet`
--
ALTER TABLE `wallet`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;

--
-- AUTO_INCREMENT for table `wallet`
--
ALTER TABLE `wallet`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `transaction`
--
ALTER TABLE `transaction`
  ADD CONSTRAINT `transaction_ibfk_1` FOREIGN KEY (`idWallet`) REFERENCES `wallet` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
